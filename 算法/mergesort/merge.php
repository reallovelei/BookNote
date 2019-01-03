<?php

class MergeClass {
    public $tmp = [];
    public function run($arr) {
        $endIndex = count($arr) - 1;
        if ($endIndex > 1) {
            $this->msort($arr, 0, $endIndex);
        }
        return $arr;
    }

    public function msort(&$arr, $startIndex, $endIndex) {
        if ($startIndex < $endIndex) {
            $midIndex = floor(($endIndex + $startIndex) / 2);
            if ($startIndex <= $midIndex)
                $this->msort($arr, $startIndex, $midIndex);
            if ($midIndex + 1 <= $endIndex)
                $this->msort($arr, $midIndex + 1, $endIndex);
            $this->merge($arr, $startIndex, $midIndex, $endIndex);
        }
    }

    // 合并有序数组
    public function merge(&$arr, $startIndex, $midIndex, $endIndex) {
        for ($t = 0; $t <= count($arr) - 1;$t++) {
            echo "{$arr[$t]} ";
        }

        echo "\nstartIndex:{$startIndex} mid:$midIndex endIndex:{$endIndex} \n";
        for ($t = $startIndex; $t <= $endIndex;$t++) {
            echo "{$arr[$t]} ";
        }
        echo "\n";
        $i = $tk = $startIndex;
        $j = $midIndex +1;
        $temp = [];

        while ($i != $midIndex + 1 && $j != $endIndex + 1) {
            echo "$i : $arr[$i]    $j : $arr[$j] \n";
            if ($arr[$i] < $arr[$j]) {
                $temp[$tk++] = $arr[$i++];
            } else {
                $temp[$tk++] = $arr[$j++];
            }
        }

        while ($i != $midIndex + 1) {
            $temp[$tk++] = $arr[$i++];
        }

        while ($j != $endIndex + 1) {
            $temp[$tk++] = $arr[$j++];
        }
      //  $arr = $temp;
        for ($i = $startIndex; $i <= $endIndex; $i++) {
            $arr[$i] = $temp[$i];
            echo "{$arr[$i]} ";
        }
        echo "\n\n\n";
        //var_dump($arr);
    }
}



$o = new MergeClass();
$arr = array(9,1,5,8,3,7,4,6,2);
$a = $o->run($arr);
var_dump($a);

