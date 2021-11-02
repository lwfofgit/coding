package Algorithm;

import java.util.LinkedHashMap;
import java.util.Map;

public class LRUCache extends LinkedHashMap {

    private int maxSize;

    public LRUCache(int maxSize){
        super(maxSize, 0.75F, true);
        this.maxSize = maxSize;
    }

    // removeEldestEntry会移除最近最少使用的元素，即：双向链表的最尾节点
    @Override
    protected boolean removeEldestEntry(Map.Entry eldest) {
        return size() > maxSize;
    }

}
