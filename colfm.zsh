# Various helpers for the zsh together with colfm.

# ccd: change the directory with colfm, quick access with C-o.
# cargs: like xargs with the current marked files
# csel: list selected files line by line

if [ -r /tmp/$USER/.colfmdir ]; then
  colfm_dir="/tmp/$USER"
else
  colfm_dir="~"
fi

ccd() {
  # This is nasty, but widgets weren't supposed to run curses apps.
  colfm.rb "$@" <$TTY
  print -n "\033[A"
  zle && zle -I                 # force redrawing of prompt
  cd "$(cat $colfm_dir/.colfmdir)"
}

cargs() {
  xargs -r -0 "$@" <$colfm_dir/.colfmsel
}

csel() {
  if tty -s; then
    tr '\0' '\n' <$colfm_dir/.colfmsel
    echo
  else
    tr '\n' '\0' |  >$colfm_dir/.colfmsel
  fi
}

zle -N ccd
bindkey "^O" ccd
