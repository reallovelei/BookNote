#include <stdio.h>

void pprint(int *a)
{
    for (int i = 0; i < 10; i++) {
        //printf("%d:%d   ", i, a[i]);
        printf("%d   ", a[i]);
    }
    printf(" \n");
}

void qsort(int *a, int left, int right)
{
    if (left >= right) {
        return ;
    }

    int i = left;
    int j = right;
    int val = a[left];

    while (i < j) {
        while (i < j && val <= a[j]) { // 如果i所在值比右侧j所在位置的值小 不用动j指针往前挪
            j--; // 继续往左找
        }

        // 如果找到比val值小的 就换到i 的位置
        a[i] = a[j];
        printf("1 i:%d   j:%d \n", i, j);

        while (i < j && val >= a[i])
        {
            i++;
        }
        a[j] = a[i];
        printf("2 i:%d   j:%d \n", i, j);
    }
    a[i] = val;

    pprint(a);

    qsort(a, left, i - 1);
    qsort(a, i + 1, right);
}

int main()
{
    int arr[10] = {6,2,3,10,1,4,7,8,9,5};
    qsort(arr, 0, 9);
    pprint(arr);
    return 0;
}

