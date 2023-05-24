<p align="center">
  <h3 align="center">Shook</h3>
<p align="center">Enable the <strong>Shell</strong> scripts could be triggered by <strong>Webhook</strong></p>
</p>

[![Release](https://img.shields.io/github/release/thankrain/shook.svg?style=for-the-badge)](https://github.com/thankrain/shook/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=for-the-badge)](/LICENSE)
[![Powered By: ThankRain](https://img.shields.io/badge/powered%20by-thankrain-red.svg?style=for-the-badge)](https://github.com/thankrain)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/thankrain/shook/release.yml?style=for-the-badge)


---

# Install

```shell
curl -sfL https://raw.githubusercontent.com/ThankRain/shook/main/install.sh | bash
```

# Usage

### Initialize Shook

```shell
$ shook init http://127.0.0.1:2399
# $ shook init http://127.0.0.1:2399 [specific_token]
# No token was set, generating token
# 52fdfc072182654f163f5f0f9a621d72
# Server initialization successfully. Please keep your token carefully.
# Shook initialization successfully!
```

### Create a new webhook

Create a new webhook to invoke hello.sh when trigger <host:port>/hello
```shell
$ shook create hello ./hello.sh
# $ shook create [hook_name] [script]
# /hello hooks created!
# $ cd D:\Develop\go\shook\src\cli ; echo $(date) > 1.txt
```

When trigger /hello, the app would cd to current folder and run the `./hello.sh` command

### Invoke the webhook

```shell
$ shook run hello
# $ shook run [hook_name]
```

**Or**

```shell
$ curl http://127.0.0.1:2399/hello
# curl <schema://host:post>/<hook_name>
```

# Notice

The server default register the path `/admin` to manage the webhooks, avoid to use this path as webhook invoke path

# Features

- [x] Basic GET and POST Webhook Without Params
- [x] Basic Operation Auth
- [ ] Support auto git repository deploy shell scripts

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
