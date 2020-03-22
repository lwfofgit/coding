package DadaStructure;

public class main {
    public static void main(String[] args) {
//        TreeNode<String> root = new TreeNode("root");
//        root.addChild("first child");
//        root.addChild("second child");
//        root.addChild("third child");
//
//        root.getChilds().get(0).addChild("first child 1");
//        root.getChilds().get(0).addChild("first child 2");
//        root.getChilds().get(2).addChild("third child 1");
//        root.getChilds().get(2).addChild("third child 2");
//
//        Order order = new Order();
//        order.post(root);

        BitMap bitMap = new BitMap(1);
        bitMap.addNum(0);
        bitMap.addNum(1);
        bitMap.addNum(6);

        bitMap.showByte(bitMap.getBytes()[0]);


    }
}
