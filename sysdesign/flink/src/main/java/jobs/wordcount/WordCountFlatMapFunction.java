package jobs.wordcount;

import org.apache.flink.api.common.functions.OpenContext;
import org.apache.flink.api.common.functions.RichFlatMapFunction;
import org.apache.flink.api.common.state.ValueState;
import org.apache.flink.api.common.state.ValueStateDescriptor;
import org.apache.flink.api.common.typeinfo.Types;
import org.apache.flink.api.java.tuple.Tuple2;
import org.apache.flink.util.Collector;

public class WordCountFlatMapFunction extends RichFlatMapFunction<String, Tuple2<String, Integer>> {
    private transient ValueState<Integer> wordCount;

    @Override
    public void open(OpenContext openContext) throws Exception {
        super.open(openContext);

        ValueStateDescriptor<Integer> desc =
                new ValueStateDescriptor<>(
                        "wordCount",
                        Types.INT
                );

        wordCount = getRuntimeContext().getState(desc);
    }

    @Override
    public void flatMap(String s, Collector<Tuple2<String, Integer>> collector) throws Exception {
        Integer count = wordCount.value();

        if (count == null) {
            count = 0;
        }

        count++;

        wordCount.update(count);
        collector.collect(new Tuple2<>(s, count));
    }
}
