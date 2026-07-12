(* [project] tnh_ai_master_template *)
(* [purpose] dna_identity: edge deployment variables *)

type execution_context = {
  platform : string;
  identity : string;
  domain   : string;
}

let current_context = {
  platform = "cloudflare_workers";
  identity = "lowercase_minimal_style";
  domain   = "arthitsiangwan.dev";
}
