import torch

# from transformers import AutoModelForCausalLM, AutoTokenizer
from transformers import LlamaForCausalLM, LlamaTokenizer

tokenizer = LlamaTokenizer.from_pretrained("./KoAlpaca/")
model = LlamaForCausalLM.from_pretrained("./KoAlpaca/")

print("KoAlpaca!")

tensor = model.generate(**tokenizer("안녕하세요?", return_tensors="pt"))

print(tensor)
