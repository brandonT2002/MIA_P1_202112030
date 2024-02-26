grammar Parser;

@header {
    import (
        commands "mia/Classes/Commands"
        interfaces "mia/Classes/Interfaces"

        "os"
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
    c1 = execute {$result = $c1.result} |
    c2 = mkdisk  {$result = $c2.result} |
    c3 = rep     {$result = $c3.result} |
    c20 = commentary {$result = $c20.result} |
    exit ;

execute returns[*commands.Execute result] locals[map[string]string params = map[string]string{}] :
    e = RW_execute (RW_path TK_equ p = TK_path {$params = map[string]string{"path": $p.text}})? {$result = commands.NewExecute($params, $e.line, $e.pos)} ;

mkdisk returns[*commands.Mkdisk result] :
    m = RW_mkdisk p = mkdiskparams {$result = commands.NewMkdisk($m.line, $m.pos, $p.result)};

mkdiskparams returns[map[string]string result] :
    {$result = map[string]string{"fit": "FF", "unit": "M"}} (p = mkdiskparam {$result[$p.result[0]] = $p.result[1]})* ;

mkdiskparam returns[[]string result] :
    RW_size TK_equ v1 = TK_number {$result = []string{"size", $v1.text}} |
    RW_fit  TK_equ v2 = TK_fit    {$result = []string{"fit",  $v2.text}} |
    RW_unit TK_equ v3 = TK_unit   {$result = []string{"unit",   $v3.text}} ;

rmdisk :
    RW_rmdisk ;

rep returns[*commands.Rep result] :
    r = RW_rep p = repparams {$result = commands.NewRep($r.line, $r.pos, $p.result)} ;

repparams returns[map[string]string result] :
    {$result = map[string]string{}} (p = repparam {$result[$p.result[0]] = $p.result[1]})* ;

repparam returns[[]string result] :
    RW_name TK_equ v1 = name    {$result = []string{"name", $v1.text}} |
    RW_path TK_equ v2 = TK_path {$result = []string{"path", $v2.text}} |
    RW_id   TK_equ v3 = TK_id   {$result = []string{"id",   $v3.text}} |
    RW_ruta TK_equ v4 = TK_path {$result = []string{"ruta", $v4.text}} ;

name returns[string result] :
    n = RW_mbr        {$result = $n.text} |
    n = RW_disk       {$result = $n.text} |
    n = RW_inode      {$result = $n.text} |
    n = RW_journaling {$result = $n.text} |
    n = RW_block      {$result = $n.text} |
    n = RW_bm_inode   {$result = $n.text} |
    n = RW_bm_block   {$result = $n.text} |
    n = RW_tree       {$result = $n.text} |
    n = RW_sb         {$result = $n.text} |
    n = RW_file       {$result = $n.text} |
    n = RW_ls         {$result = $n.text} ;

commentary returns[*commands.Commentary result] :
    c = COMMENTARY {$result = commands.NewCommentary($c.line, $c.pos, $c.text)} ;

exit : RW_exit {os.Exit(1)} ;