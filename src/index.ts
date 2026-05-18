export default {
  async fetch(request: Request, env: any, ctx: any): Promise<Response> {
    // 1. บันทึกสติลง KV เปิดสถานะสด 24 ชม.
    await env.TNH_KV.put('LAST_BOOT', new Date().toISOString());

    // 2. ดึงข้อมูลตรวจสอบความพร้อมของสมอง D1
    const { results } = await env.DB.prepare("SELECT name FROM sqlite_master LIMIT 1").all();

    // 3. เตรียมโครงสร้างดักจับ Request ยิงไปหาพอร์ต 2026 หรือฝั่งหลังบ้าน
    const url = new URL(request.url);
    
    // หน้าด่านสะพานเชื่อมโยงข้อมูล
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
          "Access-Control-Allow-Origin": "*" // เปิดทางให้คอม 200k วิ่งมาดูดของได้
        } 
      }
    );
  }
};
