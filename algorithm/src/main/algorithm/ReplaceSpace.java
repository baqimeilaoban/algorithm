package main.algorithm;

public class ReplaceSpace {
    public String replaceSpace(String str){
        int length = str.length();
        StringBuffer result = new StringBuffer();
        for (int i = 0; i < length; i++) {
            char b = str.charAt(i);
            if (String.valueOf(b).equals(" ")){
                result.append("%20");
            }else {
                result.append(b);
            }
        }
        return result.toString();
    }
}
