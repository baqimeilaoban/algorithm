package main.test;

import main.algorithm.BaseExponent;

import java.io.IOException;
import java.util.Scanner;

public class BaseExponentTest {
    public static void main(String[] args) throws IOException {
        Scanner sc = new Scanner(System.in);
        System.out.println("请输入底数：");
        double base = sc.nextDouble();
        System.out.println("输入的底数为：" + base);
        System.out.println("请输入指数：");
        int exponent = sc.nextInt();
        System.out.println("输入的指数为：" +exponent);
        BaseExponent be = new BaseExponent();
        double result = be.Power(base,exponent);
        System.out.println("数值的整数次方为：" + result);

    }
}
