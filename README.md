<h1 style="text-align: center" align="center">🛡WAFFLE🛡</h1>

[//]: # (We have a discord, come and join :)

[//]: # (https://discord.gg/kNfZzCx7se)

<div align="center">


  <img src="readme/gifs/logo.gif" alt="drawing" width="400" class="logo"/>


</div>

---

[//]: # (the problem)
<h1>❗ The problem </h1>
<div style="display: flex; justify-content: space-between">
    <p style="width: 30%">
        Imagine, that you are normal user who wants to play a game, or use a website on remote server, but something isn't right,
        the connection is so slow, and you don't know why. You are probably mad, because you can't watch your favorite
        cat videos and images. Anime is out of your reach as well, you are thinking about eating pizza from tasty <b>Da Grasso</b> 🐴
        to make your day just a little bit better and better.
    </p>
    <div align="center">
        <img src="readme/gifs/ddos_attack.gif"alt="drawing" width="500" class="logo"/>
    </div>
</div>

---

[//]: # (the solution)
<h1 style="text-align: right">💚 The solution</h1>


<div style="display: flex; justify-content: space-between">
    <div align="center">
        <img src="readme/gifs/ddos_defence.gif" alt="drawing" width="500" class="logo"/>
    </div>
    <p style="width: 30%">
        You wrote to the website administrator to use <b>Waffle</b>, the <b>open-source</b> DDOS protector, the reverse
        proxy of future, the strongest defense of all, and it's free (and not nearly done at all)! 
    </p>
</div>

## Documentation
This is github wiki's based documentation of the project (in progress):
- [Documentation](https://github.com/cebilon123/waffle/wiki/Documentation)

## Introduction
Probably you know about CloudFlare, every one knows, but this is a partially paid solution. As the open source community
we are missing a real modular and open source **Web Application Firewall** that could be used in the place of CF.
The project is huge, tho we are looking for the contributors.

## How to run / Develop ?

### Prerequisites 
+ Go 1.22+
+ golangci-lint
+ [Npcap](https://npcap.com/) (windows)
+ make (if windows, try using chocolatey)
+ openssl (if windows, try using git bash)
+ [mockery](https://vektra.github.io/mockery/latest/installation/) 

1. Create certificates and FS embed go file provider `make certs_windows`
[//]: # (2. Execute `docker compose up -d` to create needed infrastructure)
2. Set environment variables before running the proxy:

### Generate certificates
Execute make certs_windows and go through process. It should certs in the .cert directory.

### Next steps

> \[!IMPORTANT]
>
>
>
> Currently, the project can be only used with GO installed locally, but there is a plan to create usable binaries.

**How to run binaries?*
Currently there are 3 binaries: 
- collector
- proxy
- tcpproxy (partially done)

The proxy is the core binary, which can be kind of used, to run it execute:
`go run ./cmd/proxy/main.go`

Remember to add valid config to the `./cmd/proxy/config.yml` file, looks something like this:
```yaml
dns:
  - host: "google.localhost:8080"
    address: "https://google.com"
  - host: "100commitow.localhost:8080"
    address: "https://100commitow.pl"

rules:
  custom:
    - name: "request payload must contain at least one character and should contain at least one header"
      predicate: "p => LEN(p.payload) > 0 && LEN(p.headers) > 0"
    - name: "payload must be a json"
      predicate: "p => FORMAT(p.payload) == 'json'"
```

## Planned features / Architecture
To bo honest, I'm learning how to write WAF from the scratch, so this part will be updated after a while. 

- [X] XSS protection (HTML + we can take a look on sql injection)
- [ ] DDOS protection

## What I have learned?
- Neovim
- DDOS protection
- XSS /SQLI protection

# Contribution
### What do I need to know to help?
If you are looking to help to with a code contribution our project uses  **GO, k8s.** 

### How can I do that?

Never made an open source contribution before? Wondering how contributions work in the in our project? Here's a quick rundown!

Find an issue that you are interested in addressing or a feature that you would like to add.

Fork the repository associated with the issue to your local GitHub organization. This means that you will have a copy of the repository under your-GitHub-username/repository-name.

Clone the repository to your local machine using git clone. 

Create a new branch for your fix using git checkout -b branch-name-here.

Make the appropriate changes for the issue you are trying to address or the feature that you want to add.

Use git add insert-paths-of-changed-files-here to add the file contents of the changed files to the "snapshot" git uses to manage the state of the project, also known as the index.

Use git commit -m "Insert a short message of the changes made here" to store the contents of the index with a descriptive message.

Push the changes to the remote repository using git push origin branch-name-here.

Submit a pull request to the upstream repository.

Title the pull request with a short description of the changes made and the issue or bug number associated with your change. For example, you can title an issue like so "Added more log outputting to resolve #4352".

In the description of the pull request, explain the changes that you made, any issues you think exist with the pull request you made, and any questions you have for the maintainer. It's OK if your pull request is not perfect (no pull request is), the reviewer will be able to help you fix any problems and improve it!

Wait for the pull request to be reviewed by a maintainer.

Make changes to the pull request if the reviewing maintainer recommends them.

Celebrate your success after your pull request is merged! 🚀

Where can I go for help?
If you need help, you can ask questions on our Discord: https://discord.gg/33azuUWnm4

What does the Code of Conduct mean for me?

> Our Code of Conduct means that you are responsible for treating everyone on the project with respect and courtesy regardless of their identity. If you are the victim of any inappropriate behavior or comments as described in our Code of Conduct, we are here for you and will do the best to ensure that the abuser is reprimanded appropriately, per our code.

> HTML injection are attacks agains the HTML tokenization algorithm, examples:
![img.png](readme/html_injection_Samples.png)
> Basically, we need to tokenize input and check attributes, tags against a set of rules

Links:
+ [A Comprehensive Examination of Cloudflare's IP-based Distributed Denial of Service Mitigation](https://www.researchgate.net/publication/375238537_A_Comprehensive_Examination_of_Cloudflare%27s_IP-based_Distributed_Denial_of_Service_Mitigation)
+ [A Brief Study on The Evolution of Next Generation Firewall and Web Application Firewall](https://www.researchgate.net/publication/351637754_A_Brief_Study_on_The_Evolution_of_Next_Generation_Firewall_and_Web_Application_Firewall)
+ [SWAP: Mitigating XSS Attacks using a Reverse Proxy](https://sites.cs.ucsb.edu/~chris/research/doc/sess09_swap.pdf)
+ ![img.png](readme/img.png)
