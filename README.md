<p align="center">
  <h3 align="center">Shook</h3>
  <p align="center">
Shell as a Webhook  
Enable the <strong>shell</strong> scripts could be triggered by <strong>Webhook</strong></p>
</p>

[![Release](https://img.shields.io/github/release/thankrain/shook.svg?style=for-the-badge)](https://github.com/thankrain/shook/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=for-the-badge)](/LICENSE)
[![Powered By: ThankRain](https://img.shields.io/badge/powered%20by-thankrain-red.svg?style=for-the-badge)](https://github.com/thankrain)


---

### WARNING

This repo is at ***experimental*** and contains ***safety risks***, please don't use it in the production environment!!!

---

# Install

```shell
curl -sfL https://raw.githubusercontent.com/ThankRain/shook/main/install.sh | bash
```

# Usage

1. Clone this repo
2. Run `docker-compose up -d` in the repo dir
3. Create a new webhook

```shell
$ # Login
$ shook login <host:port> <username> <password>
Login Succeeded!
$ # Create a new webhook to invoke hello.sh when trigger <host:port>/hello
$ shook hello ./hello.sh
Created!
$ # When trigger /hello, the app would cd to current folder and run the `./hello.sh` command 
```

4. Invoke the webhook

```shell
shook run hello
```

Or

```shell
curl http://127.0.0.1:2399/hello
```

# Reserved

The server default register the path `/admin` to manage the webhooks, avoid to use this path as webhook invoke path

# Features

- [x] Basic GET and POST Webhook Without Params
- [ ] Basic Operation Auth

# License
```text
MIT License

Copyright (c) 2023 ThankRain

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```