import asyncio

import lmql


@lmql.query
async def hello():
    """
    argmax "Hello[WHO]" from "openai/text-ada-001" where len(WHO) < 10
    """


(await hello())[0].prompt
