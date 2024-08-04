# MakeYourCareer

## Prerequisite
Install air to run a perpetual server : [☁ Air](https://github.com/cosmtrek/air)\
Templ compiler                        : [templ](https://templ.guide/)\
NodeJs                                : [nodejs](https://nodejs.org/en)\
Make                                  : [make](https://www.gnu.org/software/make/)

## Running
NOTE: ```make docker-run``` in the root makefile does not work for now.

Install npm dependencies (tailwind) and start watching the css file:
```zsh
cd ./web
npm install
make tw-watch
```

In the project root directory run the following command to start the website server
```zsh
cd src
make dev
```
You might have to run `templ generate` before running the `make dev` command, every time you make
a change to a _*.templ*_ file. If you use neovim you can use this autocommand to run `templ 
generate` after every buffer write.
```lua
vim.api.nvim_create_autocmd({ 'BufWritePost' }, {
    desc = 'running templ generate on save',
    group = <your-group>,
    pattern = { '*.templ' },
    callback = function()
        vim.cmd(":silent !templ generate")
    end
})
```
