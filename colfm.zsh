# helper for the zsh together with colfm.

# ccd: change the directory with colfm, quick access with C-o.
ccd() {
  # This is nasty, but widgets weren't supposed to run curses apps.
  colfm "$@" <$TTY
  print -n "\033[A"
  zle && zle -I                 # force redrawing of prompt
  if [ -r /tmp/$USER/.colfmdir ]; then
    colfm_dir="/tmp/$USER"
  else
    colfm_dir="$HOME"
  fi
  cd "$(cat $colfm_dir/.colfmdir)"
}

zle -N ccd
bindkey "^O" ccd
