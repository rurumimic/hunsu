from marvin import ai_fn


@ai_fn
def classify_sentiment(tweets: list[str]) -> list[bool]:
    """
    트윗 목록이 주어지면
    각각의 트윗을 긍정(참) 또는 부정(거짓)으로 분류하고
    그에 해당하는 해당 목록을 반환합니다.
    """


if __name__ == "__main__":
    print("트윗을 분류해")
    tweets = [
        "@elonmusk: I wish the media would stop flattering me all time … it’s a bit much guys ☺️",
        "@realDonaldTrump: Please support our Capitol Police and Law Enforcement. They are truly on the side of our Country. Stay peaceful!",
        "@Sonny7: A legend of Spurs and a legend of football. Congratulations on an amazing career @GarethBale11 !! Wishing you the best luck in your next chapter mate 😁🤍",
    ]
    sentiments = classify_sentiment(tweets)
    for x, y in zip(sentiments, tweets):
        print(f"{x}: {y}")
