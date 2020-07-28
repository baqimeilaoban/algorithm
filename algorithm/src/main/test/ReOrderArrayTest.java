package main.test;

import main.algorithm.ReOrderArray;

import java.util.Scanner;

public class ReOrderArrayTest {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.println("请输入数组并以空格分隔开：");
        String str = sc.nextLine();
        String[] strs = str.trim().split("\\s{1,}");
        int[] arr = new int[strs.length];
        for (int i=0;i<strs.length;i++){
            arr[i] = Integer.parseInt(strs[i]);
        }
        ReOrderArray reOrderArray = new ReOrderArray();
        reOrderArray.reOrderArray(arr);
    }
}
