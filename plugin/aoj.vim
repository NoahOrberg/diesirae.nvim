if exists('g:loaded_aoj')
  finish
endif
let g:loaded_aoj = 1

function! s:RequireAOJ(host) abort
  return jobstart(['aoj.nvim'], { 'rpc': v:true })
endfunction


call remote#host#Register('aoj.nvim', '0', function('s:RequireAOJ'))
call remote#host#RegisterPlugin('aoj.nvim', '0', [
  \ {'type': 'command', 'name': 'AojSubmit',  'sync': 1, 'opts': {'nargs': '+'}},
  \ {'type': 'command', 'name': 'AojStatus',  'sync': 1, 'opts': {}},
  \ {'type': 'command', 'name': 'AojSelf',    'sync': 1, 'opts': {}},
  \ {'type': 'command', 'name': 'AojSession', 'sync': 1, 'opts': {}},
  \ ])
