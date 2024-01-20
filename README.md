# TODO

NextChat (ChatGPT Next Web) with Copilot

# How to run
```shell
export OPENAI_API_KEY=your_copilot_token
docker build -t test . 
docker run --name tt -p 3000:3000 -td test
```
then, open http://localhost:3000