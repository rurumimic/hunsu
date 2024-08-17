# Text Embeddings Inference

- huggingface
  - [Text Embeddings Inference](https://huggingface.co/docs/text-embeddings-inference)
  - [Supported models and hardware](https://huggingface.co/docs/text-embeddings-inference/supported_models)
- github
  - [text-embeddings-inference](https://github.com/huggingface/text-embeddings-inference)

## Build

### Install Requirements

```bash
apt install -y libssl-dev gcc
```

```bash
apt install -y protobuf-compiler
protoc --version
```

```bash
export CC=gcc
export CXX=g++
```

### Cargo Build

```bash
cargo build
```

### Run

- jinaai
  - [jina-embeddings-v2-small-en](https://huggingface.co/jinaai/jina-embeddings-v2-small-en)
  - [jina-embeddings-v2-base-en](https://huggingface.co/jinaai/jina-embeddings-v2-base-en)

```bash
model=jinaai/jina-embeddings-v2-small-en

cargo run -- --model-id $model --port 8080
```

```bash
curl 127.0.0.1:8080/embed \
    -X POST \
    -d '{"inputs":"What is Deep Learning?"}' \
    -H 'Content-Type: application/json'
```

Output:

```log
[[-0.07185282,-0.033166498,-0.007000891 ...]]
```

<details>
    <summary>Logs</summary>

```log
2024-08-17T08:05:24.500274Z  INFO text_embeddings_router: router/src/main.rs:175: Args { model_id: "jin***/****-**********-**-****l-en", revision: None, tokenization_workers: None, dtype: None, pooling: None, max_concurrent_requests: 512, max_batch_tokens: 16384, max_batch_requests: None, max_client_batch_size: 32, auto_truncate: false, default_prompt_name: None, default_prompt: None, hf_api_token: None, hostname: "0.0.0.0", port: 8080, uds_path: "/tmp/text-embeddings-inference-server", huggingface_hub_cache: None, payload_limit: 2000000, api_key: None, json_output: false, otlp_endpoint: None, otlp_service_name: "text-embeddings-inference.server", cors_allow_origin: None }
2024-08-17T08:05:24.500496Z  INFO hf_hub: ~/.cargo/registry/src/index.crates.io-6f17d22bba15001f/hf-hub-0.3.2/src/lib.rs:55: Token file not found "~/.cache/huggingface/token"
2024-08-17T08:05:24.573076Z  INFO download_pool_config: text_embeddings_core::download: core/src/download.rs:38: Downloading `1_Pooling/config.json`
2024-08-17T08:05:25.526939Z  INFO download_new_st_config: text_embeddings_core::download: core/src/download.rs:62: Downloading `config_sentence_transformers.json`
2024-08-17T08:05:25.943014Z  INFO download_artifacts: text_embeddings_core::download: core/src/download.rs:21: Starting download
2024-08-17T08:05:25.943132Z  INFO download_artifacts: text_embeddings_core::download: core/src/download.rs:23: Downloading `config.json`
2024-08-17T08:05:26.351642Z  INFO download_artifacts: text_embeddings_core::download: core/src/download.rs:26: Downloading `tokenizer.json`
2024-08-17T08:05:27.403008Z  INFO download_artifacts: text_embeddings_backend: backends/src/lib.rs:328: Downloading `model.safetensors`
2024-08-17T08:05:35.400884Z  INFO download_artifacts: text_embeddings_core::download: core/src/download.rs:32: Model artifacts downloaded in 9.457886106s
2024-08-17T08:05:35.444846Z  INFO text_embeddings_router: router/src/lib.rs:199: Maximum number of tokens per request: 8192
2024-08-17T08:05:35.445346Z  INFO text_embeddings_core::tokenization: core/src/tokenization.rs:28: Starting 10 tokenization workers
2024-08-17T08:05:35.500557Z  INFO text_embeddings_router: router/src/lib.rs:241: Starting model backend
2024-08-17T08:05:35.501282Z  INFO text_embeddings_backend_candle: backends/candle/src/lib.rs:184: Starting JinaBert model on Cpu
2024-08-17T08:07:43.809812Z  WARN text_embeddings_router: router/src/lib.rs:267: Backend does not support a batch size > 4
2024-08-17T08:07:43.809841Z  WARN text_embeddings_router: router/src/lib.rs:268: forcing `max_batch_requests=4`
2024-08-17T08:07:43.812568Z  INFO text_embeddings_router::http::server: router/src/http/server.rs:1778: Starting HTTP server: 0.0.0.0:8080
2024-08-17T08:07:43.812616Z  INFO text_embeddings_router::http::server: router/src/http/server.rs:1779: Ready
2024-08-17T08:07:53.778173Z  INFO embed{total_time="720.595901ms" tokenization_time="2.436758ms" queue_time="1.250135ms" inference_time="716.247164ms"}: text_embeddings_router::http::server: router/src/http/server.rs:706: Success
^C
2024-08-17T08:08:19.221733Z  INFO text_embeddings_router::shutdown: router/src/shutdown.rs:27: signal received, starting graceful shutdown
```

</details>

