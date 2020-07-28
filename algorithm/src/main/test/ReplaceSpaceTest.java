package main.test;

import main.algorithm.ReplaceSpace;

import java.util.Scanner;

public class ReplaceSpaceTest {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.println("请输入字符串：");
        String str = sc.nextLine();
        ReplaceSpace rs = new ReplaceSpace();
        String result = rs.replaceSpace(str);
        System.out.println(result);
    }
}
