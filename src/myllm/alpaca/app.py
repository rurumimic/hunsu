import torch
from transformers import LlamaForCausalLM, LlamaTokenizer

tokenizer = LlamaTokenizer.from_pretrained("./Alpaca-native-4bit-ggml/")
model = LlamaForCausalLM.from_pretrained("./Alpaca-native-4bit-ggml/")

print("Hello, Alpaca!")

tokens = tokenizer.encode("Hello, Alpaca!", return_tensors="pt")

output = model.generate(**tokens)

print(output)
