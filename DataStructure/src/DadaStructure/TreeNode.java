package DadaStructure;

import java.util.ArrayList;

public class TreeNode<T> {
    private T data;
    private TreeNode<T> parent;
    private ArrayList<TreeNode<T>> childs;

    public T getData() {
        return data;
    }

    public TreeNode() {
        this.data = null;
        this.parent = null;
        this.childs = new ArrayList<>();
    }

    public TreeNode(T data) {
        this.data = data;
        this.parent = null;
        this.childs = new ArrayList<>();
    }

    // 添加孩子
    public void addChild(T data) {
        this.childs.add(new TreeNode<>(data));
    }

    // 获取孩子列表
    public ArrayList<TreeNode<T>> getChilds() {
        return this.childs;
    }

    public boolean isLeafNode(TreeNode node) {
        if (node.getChilds() == null || node.getChilds().size() == 0) {
            return true;
        }
        return false;
    }

}

class Order {
    public void visit(TreeNode node) {
        System.out.println(node.getData());
    }

    // 前序遍历
    public void pre(TreeNode root) {
        if (root != null) {
            visit(root);
            ArrayList<TreeNode> childs = root.getChilds();
            for (TreeNode node : childs) {
                if (node != null) {
                    pre(node);
                }
            }
        }
    }

    // 后序遍历
    public void post(TreeNode root) {
        if (root != null) {
            ArrayList<TreeNode> childs = root.getChilds();
            for (TreeNode node : childs) {
                if (node != null) {
                    pre(node);
                }
            }
            visit(root);
        }
    }
}
