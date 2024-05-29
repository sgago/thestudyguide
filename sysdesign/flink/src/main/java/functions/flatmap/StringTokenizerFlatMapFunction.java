package functions.flatmap;

import org.apache.flink.api.common.functions.FlatMapFunction;
import org.apache.flink.api.java.tuple.Tuple2;
import org.apache.flink.util.Collector;

public class StringTokenizerFlatMapFunction implements FlatMapFunction<String, Tuple2<String, Integer>> {
    @Override
    public void flatMap(@org.jetbrains.annotations.NotNull String s, Collector<Tuple2<String, Integer>> collector) throws Exception {
        String[] words = s.toLowerCase().split("\\s+");

        for (String word : words) {
            if (!word.isEmpty()) {
                collector.collect(new Tuple2<>(word, 1));
            }
        }
    }
}
