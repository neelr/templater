
echo -n -e "\033]0;Hack Zone\007"
echo Hey Neel,
echo 
fortune -s | cowsay -f `ls /usr/local/Cellar/cowsay/3.04/share/cows | gshuf -n 1` | lolcat --seed=100 
alias ls='ls -G'
ntfs () {
	sudo /usr/local/bin/ntfs-3g /dev/$1 /Volumes/NTFS -olocal -oallow_other
}
connect () {
	ssh neelr@192.168.1.65
	expect "neelr@192.168.1.65's password: "
	send "hacker01"
}
conn () {
	ssh -l pi proxy21.rt3.io -p 33410
}
c++ () {
	g++ $1
	./a.out
}
hub () {
	cd ~/Documents/Git
}

PS1="\033[1;37m▲ \w \033[0m"
PATH=$PATH:/Users/neelredkar/anaconda3/condabin

# >>> conda initialize >>>
# !! Contents within this block are managed by 'conda init' !!
__conda_setup="$('/Users/neelredkar/anaconda3/bin/conda' 'shell.bash' 'hook' 2> /dev/null)"
if [ $? -eq 0 ]; then
    eval "$__conda_setup"
else
    if [ -f "/Users/neelredkar/anaconda3/etc/profile.d/conda.sh" ]; then
        . "/Users/neelredkar/anaconda3/etc/profile.d/conda.sh"
    else
        export PATH="/Users/neelredkar/anaconda3/bin:$PATH"
    fi
fi
unset __conda_setup
# <<< conda initialize <<<

