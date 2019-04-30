<?php
/**
 * 楼梯有n个台阶，上楼可以一步上1阶，也可以一步上2阶，一共有多少种上楼的方法？
 * 其实是一个动态规划的问题, 如果走的是1, 那么还剩下 f(n-1)种走法,如果走的是2,那么还剩下f(n-2)种走法.
 */
function getNum($n) {
    if ($n < 0) return 0;
    if ($n == 1) return 1;
    if ($n == 2) return 2;
    if ($n > 2) return getNum($n -1) + getNum($n -2);
}

echo getNum(5)."\n";
