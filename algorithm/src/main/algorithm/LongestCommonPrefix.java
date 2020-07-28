package main.algorithm;

import java.util.Arrays;

public class LongestCommonPrefix {
    public String longestCommonPrefix(String[] strs){
        if (!checkStr(strs)){
            return "";
        }
        int len = strs.length;
        StringBuilder sb = new StringBuilder();
        Arrays.sort(strs);
        int m = strs[0].length();
        int n = strs[len-1].length();
        int num = Math.min(m,n);
        for (int i = 0;i<num;i++){
            if (strs[0].charAt(i)==strs[len-1].charAt(i)){
                sb.append(strs[0].charAt(i));
            }else {
                break;
            }
        }
        return sb.toString();
    }


    public boolean checkStr(String[] strs){
        boolean flag = false;
        if (strs!=null){
            //遍历strs检查元素值
            for (int i=0;i<strs.length;i++){
                if (strs[i]!=null && strs[i].length()!=0){
                    flag = true;
                }else {
                    flag = false;
                    break;
                }
            }
        }
        return false;
    }
}
