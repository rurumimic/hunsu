from marvin import ai_fn


@ai_fn
def list_fruits(n: int) -> list[str]:
    """과일 이름 N개 목록을 만들어"""


if __name__ == "__main__":
    print("과일 이름 3개 목록을 만들어")
    fruits = list_fruits(3)
    print(fruits)
