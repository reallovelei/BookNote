<?php

$arr = [6,2,3,10,1,4,7,8,9,5];

function qs(&$arr, $left, $right)
{
    if ($left >= $right) {
        return ;
    }

    $i = $left;
    $j = $right;
    $val = $arr[$left];

    while ($i < $j) {
        while ($i < $j && $val < $arr[$j])
        {
            $j--; // 往左走
        }
        $arr[$i] = $arr[$j];

        while ($i < $j && $val >= $arr[$i])
        {
            $i++; // 往右走
        }
        $arr[$j] = $arr[$i];
    }
    $arr[$i] = $val;
    qs($arr, $left, $i -1 );
    qs($arr, $i + 1, $right );
}
//qs($arr, 0 , 9);

function maopao($arr)
{
    $cnt = count($arr);
    for ($i = 0; $i < $cnt; $i++) {
        for ($j = $cnt - 1; $j > $i; $j--)
        {
            if ($arr[$j - 1] > $arr[$j]) {
                $x = $arr[$j];
                $arr[$j] = $arr[$j - 1];
                $arr[$j - 1] = $x;
            }
        }
    }
    return $arr;
}
$arr = maopao($arr);
var_dump($arr);
//var_dump($arr);
