" credits: http://ornicar.github.com/2011/02/12/ranger-as-vim-file-manager.html
" Use colfm as vim file manager (https://github.com/lian/colfm.git)
function! Colfm()
    " Get a temp file name without creating it
    let tmpfile = substitute(system('mktemp -u'), '\n', '', '')
    " Launch colfm, passing it the temp file name
    silent exec '!COLFM_RETURN_FILE='.tmpfile.' colfm.rb -'
    " If the temp file has been written by colfm
    if filereadable(tmpfile)
        " Get the selected file name from the temp file
        let filetoedit = system('cat '.tmpfile)
        exec 'edit '.filetoedit
        call delete(tmpfile)
    endif
    redraw!
endfunction

map <C-x> :call Colfm()<cr>
