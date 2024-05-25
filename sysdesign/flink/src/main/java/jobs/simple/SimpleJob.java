package jobs.example;

import org.apache.flink.api.common.functions.FilterFunction;
import org.apache.flink.configuration.Configuration;
import org.apache.flink.streaming.api.datastream.DataStream;
import org.apache.flink.streaming.api.environment.StreamExecutionEnvironment;

public class SimpleJob {
    public static void main(String[] args) throws Exception {
        Configuration conf = new Configuration();
        StreamExecutionEnvironment env =
                StreamExecutionEnvironment.createLocalEnvironmentWithWebUI(conf);

        DataStream<Person> flintstones = env.fromData(
                new Person("Fred", 35),
                new Person("Wilma", 35),
                new Person("Pebbles", 2),
                new Person("Whatever", 99));

        DataStream<Person> adults = flintstones.filter(new FilterFunction<Person>() {
            @Override
            public boolean filter(Person person) {
                return person.age >= 18;
            }
        });

        adults.print();

        env.execute();
    }

    public static class Person {
        public String name;
        public Integer age;
        public Person() {}

        public Person(String name, Integer age) {
            this.name = name;
            this.age = age;
        }

        public String toString() {
            return this.name + ": age " + this.age.toString();
        }
    }
}

