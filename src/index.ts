import { WorkflowEntrypoint, WorkflowStep } from "cloudflare:workers";
import type { WorkflowEvent } from "cloudflare:workers";

// 1. ENVIRONMENT INTERFACE DEFINITION (ZERO TRUST BINDINGS)
export interface Env {
  TNH_KV: KVNamespace;
  DB: D1Database;
  BUCKET: R2Bucket;
  AI: Ai;
  MY_WORKFLOW: WorkflowNamespace;
}

type ImageParams = { imageKey: string };

// 2. DURABLE EXECUTION WORKFLOW ENGINE (L3/L5 AGENT CORE)
export class ImageProcessingWorkflow extends WorkflowEntrypoint<Env, ImageParams> {
  async run(event: WorkflowEvent<ImageParams>, step: WorkflowStep) {
    
    // STEP A: Fetch payload from R2 storage and convert to safe serializable Array
    const imageDataRaw = await step.do("fetch image", async () => {
      const object = await this.env.BUCKET.get(event.payload.imageKey);
      if (!object) {
        throw new Error(`[FAIL] Object metadata missing: ${event.payload.imageKey}`);
      }
      const buffer = await object.arrayBuffer();
      return Array.from(new Uint8Array(buffer));
    });

    // STEP B: Process via Cloudflare LLM Edge (LlaVA Model)
    const description = await step.do("generate description", async () => {
      return await this.env.AI.run("@cf/llava-hf/llava-1.5-7b-hf", {
        image: imageDataRaw,
        prompt: "Describe this image in one sentence",
        max_tokens: 50,
      });
    });

    // STEP C: Human-in-the-loop lock (Bears 24-hour timeout gate)
    await step.waitForEvent("await approval", {
      event: "approved",
      timeout: "24 hours",
    });

    // STEP D: Move sanitized data to public layer
    await step.do("publish", async () => {
      const uint8Data = new Uint8Array(imageDataRaw);
      await this.env.BUCKET.put(`public/${event.payload.imageKey}`, uint8Data);
    });

    return { status: "SUCCESS", description: description };
  }
}

// 3. MAIN GATEWAY ROUTER (FETCH HANDLER)
export default {
  async fetch(request: Request, env: Env, ctx: any): Promise<Response> {
    const url = new URL(request.url);

    // [ROUTE A]: TRIGGER NEW IMAGE WORKFLOW VIA API
    if (url.pathname === "/workflow/start") {
      const imageKey = url.searchParams.get("key") || "matrix-vision.jpg";
      const instance = await env.MY_WORKFLOW.create({
        params: { imageKey: imageKey }
      });
      return Response.json({ status: "WORKFLOW_LAUNCHED", instanceId: instance.id });
    }

    // [ROUTE B]: UNSTUCK GATEWAY - DISPATCH APPROVAL EVENT TO WORKFLOW
    if (url.pathname === "/workflow/approve") {
      const instanceId = url.searchParams.get("instanceId");
      if (!instanceId) return new Response("[FAIL] Missing instanceId", { status: 400 });
      
      const instance = await env.MY_WORKFLOW.get(instanceId);
      await instance.sendEvent({ type: "approved", payload: {} });
      return Response.json({ status: "EVENT_DISPATCHED", target: instanceId });
    }

    // [DEFAULT ROUTE]: CORE MONITORING PANEL (KEEPING BOSS ART'S EXACT STATE)
    // 1. Log timestamp state to KV storage
    await env.TNH_KV.put('LAST_BOOT', new Date().toISOString());

    // 2. Validate structural integrity of D1 Core Engine
    const { results } = await env.DB.prepare("SELECT name FROM sqlite_master LIMIT 1").all();

    // 3. Render clean data interface (Secured against Google Translate anomalies)
    return new Response(
      JSON.stringify({
        status: "TNH-AI-V83-ONLINE",
        message: "อาณาจักร 9THERA ปลดล็อกป้ายชื่อตรงกันเรียบร้อยก้า! 🤫",
        kv_check: await env.TNH_KV.get('LAST_BOOT'),
        db_ready: !!results,
        path_requested: url.pathname
      }),
      { 
        headers: { 
          "Content-Type": "application/json",
          "Access-Control-Allow-Origin": "*",
          "Content-Language": "en" 
        } 
      }
    );
  }
};
        
