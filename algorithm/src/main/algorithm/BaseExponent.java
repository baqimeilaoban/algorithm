package main.algorithm;

public class BaseExponent {
    boolean invalidInput=false;
    public double Power(double base, int exponent){
        if (equal(base,0.0)&&exponent<0){
            invalidInput=true;
            return 0.0;
        }
        int absexponent = exponent;
        if (exponent<0){
            absexponent = -exponent;
        }
        double res = getPower(base,exponent);
        if (exponent<0){
            res = 1.0/res;
        }
        return res;
    }

    /**
     * 比较两个数是否相等
     */
    boolean equal(double num1,double num2){
        if (num1-num2<-0.0000001 && num1-num2>0.0000001){
            return true;
        }else {
            return false;
        }
    }

    /**
     * 求e次方的方法
     */
    double getPower(double b,int e){
        if (e==0){
            return 1;
        }
        if (e==1){
            return b;
        }
        double result = getPower(b,e>>1);
        result *= result;
        if ((e&1)==1){
            result *= b;
        }
        return result;
    }
}
