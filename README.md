我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

# templater

<p align="center">
    <a href="https://templaterx.now.sh" alt="Templates">
        <img src="https://img.shields.io/endpoint?url=https://templater-api.hacker22.repl.co/api/badges/templates&label=Templates" /></a>
  <a href="https://templaterx.now.sh" alt="Server">
        <img src="https://img.shields.io/endpoint?url=https://templater-api.hacker22.repl.co/api/badges&label=Server%20Status" /></a>
  <a href="https://opensource.org/licenses/MIT" alt="LICENSE">
        <img src="https://badgen.net/github/license/neelr/templater" /></a>
    <a href="https://github.com/neelr/templater/commits/master" alt="Commits">
        <img src="https://badgen.net/github/last-commit/neelr/templater" /></a>
    <a href="https://github.com/neelr/templater/issues" alt="Closed Issues">
        <img src="https://badgen.net/github/closed-issues/neelr/templater" /></a>
     <a href="https://github.com/neelr/templater/issues" alt="Open Issues">
        <img src="https://badgen.net/github/open-issues/neelr/templater" /></a>
      <a href="https://github.com/neelr/templater/actions" alt="Actions">
        <img src="https://badgen.net/github/checks/neelr/templater" /></a>
  <a href="https://github.com/neelr/templater/releases" alt="Release">
        <img src="https://badgen.net/github/release/neelr/templater" /></a>
  
</p>

A cli to create and share code templates and structure! This makes life super easy, so you don't have to create the React folder structure, and can create production ready structure in seconds! Also the best part about this, is since it has integration with GitHub, you can upload and download other people's structures, so the possibilities are endless!

### Installation

1. `go get -u github.com/neelr/templater/cmd/plate` Get the package from github

2. `go install github.com/neelr/templater/cmd/plate` Install the package!

### Docs

##### `plate create name`

Creates a template using the files in the current directory, and calls it "name"!

##### `plate load name`

Loads the specified temnplate into the current directory

##### `plate delete name`

Deletes the template "name"

##### `plate list`

Lists all installed templates

##### `plate login`

Logs you in so you can upload your templates so others can download them!

##### `plate upload name`

After you've logged in, you can upload a plate so people can see it! It'll give you a slug in the format `githubUsername/templateName` which others can use to download your template!

##### `plate deletefromserver name`

Delete the template you uploaded in the past!

##### `plate get githubUsername/templateName`

Install the template someone uploaded! You can load it anytime you want!

## TODO:

-   [x] Create a Nextjs UI so people can view people's templates
-   [x] Create a landing page ~along with a logo~
-   [x] Add Badges

Open to contributers! Just open up an issue/PR!
