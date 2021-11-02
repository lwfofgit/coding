package DadaStructure;

/*
 1 bitmap是一个为数组，用byte[]来构建；下标是具体的值，值0、1表示具体的值是否出现
 2 核心方法 getIndex: 获取在byte[]里下标 getPosition: 获取byte里的位
 3 可以用了大数据量的去重 排序和简单统计「例如：统计网站用户一年的登录次数」
 */

public class BitMap {

    private byte[] bytes;
    private int Size;

    public BitMap(int size){
        this.bytes = new byte[size];
        this.Size = size;
    }

    public byte[] getBytes() {
        return bytes;
    }

    public void setBytes(byte[] bytes) {
        this.bytes = bytes;
    }

    public int getSize() {
        return Size;
    }

    public void setSize(int size) {
        Size = size;
    }

    /*
	获取在字节数组中下标，num除以8 就得到num在数组中的位置
	num / 8 == num >> 3
    */
    public int getIndex(int num){
        return num >> 3;
    }

    /*
        获取在一个字节中的位的位置，byte[index]中的第几位
        num % 8得到在一个字节中的位置
    */
    public int getPosition(int num){
        return num & (8 -1);
    }

    /*
	标记指定数字（num）在bitmap中的值，标记其已经出现过
	将1左移position后，那个位置自然就是1，然后和以前的数据做或运算，这样，那个位置就替换成1了
    */
    public void  addNum(int num){
        this.bytes[getIndex(num)] |= 1 << getPosition(num);
    }

    /*
	判断num是否在位图中
	将1左移position位置后，那个位置自然就是1，然后和以前的数据做与运算，判断是否为1，入
	如果结果为1，则以前那个位置就是1，否则以前的那个位置的数是0
    */
    public boolean contians(int num){
        return (this.bytes[getIndex(num)] &= 1 << getPosition(num)) == 1;
    }


    /**
     * 打印byte类型的变量<br/>
     * 将byte转换为一个长度为8的byte数组，数组每个值代表bit
     */
    public void showByte(byte b) {
        byte[] array = new byte[8];
        for (int i = 0; i < 8; i++) {
            array[i] = (byte) (b & 1);
            b = (byte) (b >> 1);
        }

        for (byte b1 : array) {
            System.out.print(b1);
            System.out.print(" ");
        }

        System.out.println();


    }
}
