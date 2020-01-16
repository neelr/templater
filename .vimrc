" Call the .vimrc.plug file
let g:airline_powerline_fonts = 1
set guifont=Hurmit\ Nerd\ Font:h11
set encoding=UTF-8
if filereadable(expand("~/.vimrc.plug"))
     source ~/.vimrc.plug
 endif
 " System specific configuration
runtime custom.vim


" noremap <up> <nop>
" noremap <up> <nop>
" noremap <down> <nop>
" noremap <down> <nop>
" noremap <left> <nop>
" noremap <right> <nop>
" noremap <left> <nop>
" noremap <right> <nop>
" B-A-<start>

" Conditional plugin excludes
if !exists('g:pathogen_disabled')
	let g:pathogen_disabled = []
endif
if !has('python3')
	call extend(g:pathogen_disabled, ['vimsence', 'YouCompleteMe'])
endif
if has('win32') || has('win64')
	call extend(g:pathogen_disabled, ['cursorword'])
endif
 set bs=2
 set number
 inoremap {<CR> {<CR>}<C-o>O
 inoremap ( ()<Left>
inoremap  "  ""<Left>
inoremap  '  ''<Left>
inoremap [<CR> [<CR>]<C-o>O

" Load ./ftplugin/*, ./indent/* and tries to detect filetype
" cf. https://vi.stackexchange.com/a/10125
filetype plugin indent on

" Enable modelines
" cf. http://vim.wikia.com/wiki/Modeline_magic
set modeline
set modelines=5

" Remove annoying key sequences delays
set ttimeoutlen=100

" Avoid default behavior of selecting first entry in completion menu
set completeopt=longest,menuone

" Default spell checking language to English
set spelllang=en
" Store custom spell checking dictionary in vim folder
set spellfile=~/.vim/dictionary.utf-8.add

" Persistent undo
if has('persistent_undo')
	let undo_path = $HOME . '/.local/share/vim/undo'
	call system('mkdir -p ' . undo_path)
	let &undodir = undo_path
	set undofile
endif

" Indentation set to use smart tabs
" cf. http://vim.wikia.com/wiki/Indent_with_tabs,_align_with_spaces
set noexpandtab
set copyindent
set preserveindent
set softtabstop=0
set shiftwidth=8
set tabstop=8

set mouse=a

set number
set numberwidth=1

" Display invisible characters. Tabs and trailing space.
set list
set listchars=tab:\·\ ,trail:.,extends:#,nbsp:.

syntax on
