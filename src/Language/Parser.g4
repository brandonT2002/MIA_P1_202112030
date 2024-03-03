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

commands returns[[]interfaces.Command result] :
    l = commands c = command {$result = $l.result;; $result = append($result, $c.result)} |
    c = command              {$result = []interfaces.Command{$c.result}                 } ;

command returns[interfaces.Command result] :
    c1 = execute {$result = $c1.result}      |
    c2 = mkdisk  {$result = $c2.result}      |
    c3 = rmdisk  {$result = $c3.result}      |
    c4 = fdisk   {$result = $c4.result}      |
    c20 = rep    {$result = $c20.result}     |
    c21 = commentary {$result = $c21.result} |
    exit ;

execute returns[*commands.Execute result] :
    e = RW_execute RW_path TK_equ p = TK_path {$result = commands.NewExecute($e.line, $e.pos, map[string]string{"path": $p.text})} |
    e = RW_execute                            {$result = commands.NewExecute($e.line, $e.pos, map[string]string{})               } ;

mkdisk returns[*commands.Mkdisk result] :
    m = RW_mkdisk p = mkdiskparams {$result = commands.NewMkdisk($m.line, $m.pos, $p.result)                                  } |
    m = RW_mkdisk                  {$result = commands.NewMkdisk($m.line, $m.pos, map[string]string{"fit": "FF", "unit": "M"})} ;

mkdiskparams returns[map[string]string result] :
    l = mkdiskparams p = mkdiskparam {$result = $l.result;; $result[$p.result[0]] = $p.result[1]                       } |
    p = mkdiskparam                  {$result = map[string]string{"fit": "FF", "unit": "M", $p.result[0]: $p.result[1]}};

mkdiskparam returns[[]string result] :
    RW_size TK_equ v1 = TK_number {$result = []string{"size", $v1.text}} |
    RW_fit  TK_equ v2 = TK_fit    {$result = []string{"fit",  $v2.text}} |
    RW_unit TK_equ v3 = TK_unit   {$result = []string{"unit", $v3.text}} ;

rmdisk returns[*commands.Rmdisk result] :
    r = RW_rmdisk RW_driveletter TK_equ p = TK_id {$result = commands.NewRmdisk($r.line, $r.pos, map[string]string{"driveletter": $p.text})} |
    r = RW_rmdisk                                 {$result = commands.NewRmdisk($r.line, $r.pos, map[string]string{})                      } ;

fdisk returns[*commands.Fdisk result] :
    f = RW_fdisk p = fdiskparams {$result = commands.NewFdisk($f.line, $f.pos, $p.result)} |
    f = RW_fdisk                 {$result = commands.NewFdisk($f.line, $f.pos, map[string]string{"unit": "K", "fit": "WF", "type": "P"})} ;

fdiskparams returns[map[string]string result] :
    l = fdiskparams p = fdiskparam {$result = $l.result;; $result[$p.result[0]] = $p.result[1]} |
    p = fdiskparam                 {$result = map[string]string{"unit": "K", "fit": "WF", "type": "P", $p.result[0]: $p.result[1]}} ;

fdiskparam returns[[]string result] :
    RW_size        TK_equ v1 = TK_number {$result = []string{"size",        $v1.text}} |
    RW_driveletter TK_equ v2 = TK_id     {$result = []string{"driveletter", $v2.text}} |
    RW_name        TK_equ v3 = TK_id     {$result = []string{"name",        $v3.text}} |
    RW_unit        TK_equ v4 = TK_unit   {$result = []string{"unit",        $v4.text}} |
    RW_type        TK_equ v5 = TK_type   {$result = []string{"type",        $v5.text}} |
    RW_fit         TK_equ v6 = TK_fit    {$result = []string{"fit",         $v6.text}} |
    RW_delete      TK_equ v7 = RW_full   {$result = []string{"delete",      $v7.text}} |
    RW_add         TK_equ v8 = TK_number {$result = []string{"add",         $v8.text}} ;

rep returns[*commands.Rep result] :
    r = RW_rep p = repparams {$result = commands.NewRep($r.line, $r.pos, $p.result)          } |
    r = RW_rep               {$result = commands.NewRep($r.line, $r.pos, map[string]string{})} ;

repparams returns[map[string]string result] :
    l = repparams p = repparam {$result = $l.result;; $result[$p.result[0]] = $p.result[1]} |
    p = repparam               {$result = map[string]string{$p.result[0]: $p.result[1]}   } ;

repparam returns[[]string result] :
    RW_name TK_equ v1 = name    {$result = []string{"name", $v1.text}} |
    RW_path TK_equ v2 = TK_path {$result = []string{"path", $v2.text}} |
    RW_ruta TK_equ v4 = TK_path {$result = []string{"ruta", $v4.text}} |
    RW_id   TK_equ v3 = TK_id   {$result = []string{"id",   $v3.text}} ;

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