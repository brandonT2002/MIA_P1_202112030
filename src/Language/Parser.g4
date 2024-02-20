grammar Parser;

@header {
    import (
        commands "mia/Classes/Commands"
        interfaces "mia/Classes/Interfaces"
    )
}

options {
    language = Go;
    tokenVocab = Scanner;
}

init returns[[]interfaces.Command result] :
    c = commands EOF {$result = $c.result             } |
    EOF              {$result = []interfaces.Command{}} ;

commands returns[[]interfaces.Command result] locals[[]interfaces.Command cmds = []interfaces.Command{}] :
    (c = command {$cmds = append($cmds, $c.result)})* {$result = $cmds} ;

command returns[interfaces.Command result] :
    c = execute {$result = $c.result} ;

execute returns[*commands.Execute result] locals[map[string]string params = map[string]string{}] :
    e = RW_execute (RW_path TK_equ p = TK_path {$params["path"] = $p.text})? {$result = commands.NewExecute($params, $e.line, $e.pos)} ;