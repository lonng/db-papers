# Database Papers

This is a comprehensive list of papers on database theory for understanding and building database systems. It covers various aspects of database systems, including the essential theoretical background, classic system design, and multiple modules within the database.

The list is organized into different categories and subcategories for easy navigation. Each paper is accompanied by a title, author, and publication year, along with a link to the full text if available.

This collection serves as a learning and training resource primarily for the Tencent Cloud Database Team and is also open to external researchers, students, and learners interested in database systems.

**In case you are reading this and making the effort to comprehend these papers, we would really like to have a conversation with you regarding opportunities at Tencent Cloud Database Team.**

## Contribution

This list is generated from a Sheet document automatically. If you have any suggestions or would like to contribute to this list, please feel free to file an issue. And we will update our sheet to make the chagnes available for public.

Any contribution that can help improve this list and make it more comprehensive and useful to the community are welcome. Here are some ways you can contribute:

- **Add a new paper**: If you have a paper that you think should be included in this list, please file an issue to provide the paper's title, author, publication year, and a link to the full text (if available).
- **Update an existing paper**: If you find any errors or outdated information in the list, please file an issue to provide the correct information.
- **Remove a paper**: If you think a paper is no longer relevant or useful, please file an issue to suggest its removal.
- **General suggestions**: If you have any general suggestions or feedback on how to improve this list, please file an issue to share your thoughts.

## Table of Contents

- [Basics](#basics)
  - [Essentials](#essentials)
  - [Consensus](#consensus)
  - [Consistency](#consistency)
  - [Workload](#workload)
  - [Network](#network)
  - [Quality](#quality)
- [System Design](#system-design)
  - [RDBMS](#rdbms)
  - [NoSQL](#nosql)
- [SQL Engine](#sql-engine)
  - [Optimizer Framework](#optimizer-framework)
  - [Transformation](#transformation)
  - [Nested Query](#nested-query)
  - [Functional Dependencies](#functional-dependencies)
  - [Join Order](#join-order)
  - [Cost Model](#cost-model)
  - [Statistics](#statistics)
  - [Probabilistic Counting](#probabilistic-counting)
  - [Execution Engine](#execution-engine)
  - [MPP Optimizations](#mpp-optimizations)
- [Storage Engine](#storage-engine)
  - [Storage Structure](#storage-structure)
  - [Transaction](#transaction)
  - [Scheduling](#scheduling)
## Basics

### Essentials

- [A relational model of data for large shared data banks](papers/essentials/a-relational-model-of-data-for-large-shared-data-banks.pdf) (1970) - Codd, Edgar F.
- [SEQUEL: A structured English query language](papers/essentials/sequel:-a-structured-english-query-language.pdf) (1974) - Chamberlin, Donald D., and Raymond F. Boyce.
- [INGRES: a relational data base system](papers/essentials/ingres:-a-relational-data-base-system.pdf) (1975) - Held, G. D., M. R. Stonebraker, and Eugene Wong.
- [Extending the database relational model to capture more meaning](papers/essentials/extending-the-database-relational-model-to-capture-more-meaning.pdf) (1979) - Codd, Edgar F. .
- [A critique of the SQL database language](papers/essentials/a-critique-of-the-sql-database-language.pdf) (1984) - Date, C. J.
### Consensus

- [The Part-Time Parliament](papers/consensus/the-part-time-parliament.pdf) (1998) - Lamport, Leslie.
- [Paxos Made Simple](papers/consensus/paxos-made-simple.pdf) (2001) - Lamport, Leslie.
- [Consensus: Bridging theory and practice](papers/consensus/consensus:-bridging-theory-and-practice.pdf) (2014) - Ongaro, Diego.
- [In search of an understandable consensus algorithm (extended version)](papers/consensus/in-search-of-an-understandable-consensus-algorithm-(extended-version).pdf) (2014) - Ongaro, Diego, and John Ousterhout.
- [Distributed consensus revised](papers/consensus/distributed-consensus-revised.pdf) (2019) - Howard, Heidi.
- [A Generalised Solution to Distributed Consensus](papers/consensus/a-generalised-solution-to-distributed-consensus.pdf) (2019) - Howard, Heidi, and Richard Mortier. .
- [Paxos vs Raft: Have we reached consensus on distributed consensus?](papers/consensus/paxos-vs-raft:-have-we-reached-consensus-on-distributed-consensus?.pdf) (2020) - Howard, Heidi, and Richard Mortier.
### Consistency

- [Consistency Tradeoffs in Modern Distributed Database System Design](papers/consistency/consistency-tradeoffs-in-modern-distributed-database-system-design.pdf) (2012) - Abadi, Daniel.
- [Logical physical clocks and consistent snapshots in globally distributed databases](papers/consistency/logical-physical-clocks-and-consistent-snapshots-in-globally-distributed-databases.pdf) (2014) - Kulkarni S S, Demirbas M, Madappa D, et al.
- [Ark: A Real-World Consensus Implementation](papers/consistency/ark:-a-real-world-consensus-implementation.pdf) (2014) - Kasheff, Zardosht, and Leif Walsh.
- [PolarFS: an ultra-low latency and failure resilient distributed file system for shared storage cloud database](papers/consistency/polarfs:-an-ultra-low-latency-and-failure-resilient-distributed-file-system-for-shared-storage-cloud-database.pdf) (2018) - Cao, Wei, et al.
- [Anna: A kvs for any scale](papers/consistency/anna:-a-kvs-for-any-scale.pdf) (2018) - Wu, Chenggang, et al.
- [Strong and efficient consistency with consistency-aware durability](papers/consistency/strong-and-efficient-consistency-with-consistency-aware-durability.pdf) (2021) - Ganesan, Aishwarya, et al. .
### Workload

- [TPC-H Analyzed: Hidden Messages and Lessons Learned from an Influential Benchmark](papers/workload/tpc-h-analyzed:-hidden-messages-and-lessons-learned-from-an-influential-benchmark.pdf) (2013) - Boncz, Peter, Thomas Neumann, and Orri Erling.
- [Quantifying TPCH Choke Points and Their Optimizations](papers/workload/quantifying-tpch-choke-points-and-their-optimizations.pdf) (2020) - Dreseler, Markus, et al.
### Network

- [The End of Slow Networks: It's Time for a Redesign](papers/network/the-end-of-slow-networks:-it's-time-for-a-redesign.pdf) (2015) - Binnig, Carsten, et al.
- [Accelerating Relational Databases by Leveraging Remote Memory and RDMA](papers/network/accelerating-relational-databases-by-leveraging-remote-memory-and-rdma.pdf) (2016) - Li, Feng, et al.
- [Don't Hold My Data Hostage: A Case for Client Protocol Redesign](papers/network/don't-hold-my-data-hostage:-a-case-for-client-protocol-redesign.pdf) (2017) - Raasveldt, Mark, and Hannes Mühleisen.
### Quality

- [Testing the Accuracy of Query Optimizers](papers/quality/testing-the-accuracy-of-query-optimizers.pdf) (2012) - Gu, Zhongxian, Mohamed A. Soliman, and Florian M.
## System Design

### RDBMS

- [System R: Relational Approach to Database Management](papers/rdbms/system-r:-relational-approach-to-database-management.pdf) (1976) - Astrahan, Morton M., et al.
- [The design and implementation of INGRES](papers/rdbms/the-design-and-implementation-of-ingres.pdf) (1976) - Stonebraker, Michael, et al.
- [The design of Postgres](papers/rdbms/the-design-of-postgres.pdf) (1986) - Stonebraker, Michael, and Lawrence A. Rowe.
- [Query Processing in Main Memory Database Management Systems](papers/rdbms/query-processing-in-main-memory-database-management-systems.pdf) (1986) - Lehman, Tobin J., and Michael J. Carey.
- [Megastore: Providing Scalable, Highly Available Storage for Interactive Services](papers/rdbms/megastore:-providing-scalable,-highly-available-storage-for-interactive-services.pdf) (2011) - Baker J, Bond C, Corbett J C, et al.
- [Spanner: Google's globally distributed database](papers/rdbms/spanner:-google's-globally-distributed-database.pdf) (2013) - Corbett, James C., et al.
- [Online, Asynchronous Schema Change in F1](papers/rdbms/online,-asynchronous-schema-change-in-f1.pdf) (2013) - Rae, Ian, et al.
- [Amazon aurora: Design considerations for high throughput cloud-native relational databases](papers/rdbms/amazon-aurora:-design-considerations-for-high-throughput-cloud-native-relational-databases.pdf) (2017) - Verbitski, Alexandre, et al.
- [Looking Back at Postgres](papers/rdbms/looking-back-at-postgres.pdf) (2019) - Hellerstein, Joseph M.
- [CockroachDB: The Resilient Geo-Distributed SQL Database](papers/rdbms/cockroachdb:-the-resilient-geo-distributed-sql-database.pdf) (2020) - Taft, Rebecca, et al.
- [F1 Lightning: HTAP as a Service](papers/rdbms/f1-lightning:-htap-as-a-service.pdf) (2020) - Yang, Jiacheng, et al. .
- [TiDB: a Raft-based HTAP database](papers/rdbms/tidb:-a-raft-based-htap-database.pdf) (2020) - Huang, Dongxu, et al.
- [PolarDB Serverless: A Cloud Native Database for Disaggregated Data Centers](papers/rdbms/polardb-serverless:-a-cloud-native-database-for-disaggregated-data-centers.pdf) (2021) - Cao, Wei, et al.
### NoSQL

- [Bigtable: A Distributed Storage System for Structured Data](papers/nosql/bigtable:-a-distributed-storage-system-for-structured-data.pdf) (2006) - Chang, Fay, et al.
- [Dynamo: Amazon’s Highly Available Key-value Store](papers/nosql/dynamo:-amazon’s-highly-available-key-value-store.pdf) (2007) - DeCandia, Giuseppe, et al.
- [PNUTS: Yahoo!’s Hosted Data Serving Platform](papers/nosql/pnuts:-yahoo!’s-hosted-data-serving-platform.pdf) (2008) - Cooper, Brian F., et al.
- [Cassandra - A Decentralized Structured Storage System](papers/nosql/cassandra---a-decentralized-structured-storage-system.pdf) (2010) - Lakshman, Avinash, and Prashant Malik.
- [Windows azure storage: a highly available cloud storage service with strong consistency](papers/nosql/windows-azure-storage:-a-highly-available-cloud-storage-service-with-strong-consistency.pdf) (2011) - Calder, Brad, et al.
- [Azure data lake store: a hyperscale distributed file service for big data analytics](papers/nosql/azure-data-lake-store:-a-hyperscale-distributed-file-service-for-big-data-analytics.pdf) (2017) - Ramakrishnan, Raghu, et al.
- [PNUTS to Sherpa: Lessons from Yahoo!’s Cloud Database](papers/nosql/pnuts-to-sherpa:-lessons-from-yahoo!’s-cloud-database.pdf) (2019) - Cooper, Brian F., et al.
## SQL Engine

### Optimizer Framework

- [Access Path Selection in a Relational Database Management System](papers/optimizer-framework/access-path-selection-in-a-relational-database-management-system.pdf) (1979) - Selinger, P. Griffiths, et al.
- [Query Optimization by Simulated Annealing](papers/optimizer-framework/query-optimization-by-simulated-annealing.pdf) (1987) - Ioannidis, Yannis E., and Eugene Wong.
- [The EXODUS Optimizer Generator](papers/optimizer-framework/the-exodus-optimizer-generator.pdf) (1987) - Graefe, Goetz, and David J. DeWitt.
- [Extensible/Rule Based Query Rewrite Optimization in Starburst](papers/optimizer-framework/extensible-rule-based-query-rewrite-optimization-in-starburst.pdf) (1992) - Pirahesh, Hamid, Joseph M. Hellerstein, and Waqar Hasan. .
- [The Volcano Optimizer Generator- Extensibility and Efficient Search](papers/optimizer-framework/the-volcano-optimizer-generator--extensibility-and-efficient-search.pdf) (1993) - Graefe, Goetz, and William J. McKenna.
- [The Cascades Framework for Query Optimization](papers/optimizer-framework/the-cascades-framework-for-query-optimization.pdf) (1995) - Graefe, Goetz.
- [ An Overview of Query Optimization in Relational Systems](papers/optimizer-framework/-an-overview-of-query-optimization-in-relational-systems.pdf) (1998) - Chaudhuri, Surajit.
- [LEO – DB2’s LEarning Optimizer](papers/optimizer-framework/leo-–-db2’s-learning-optimizer.pdf) (2001) - Stillger, Michael, et al.
- [Robust Query Processing through Progressive Optimization](papers/optimizer-framework/robust-query-processing-through-progressive-optimization.pdf) (2004) - Markl, Volker, et al.
- [Orca: A Modular Query Optimizer Architecture for Big Data](papers/optimizer-framework/orca:-a-modular-query-optimizer-architecture-for-big-data.pdf) (2014) - Soliman, Mohamed A., et al.
- [Parallelizing Query Optimization on Shared-Nothing Architectures](papers/optimizer-framework/parallelizing-query-optimization-on-shared-nothing-architectures.pdf) (2015) - Trummer, Immanuel, and Christoph Koch.
- [The MemSQL Query Optimizer: A modern optimizer for real-time analytics in a distributed database](papers/optimizer-framework/the-memsql-query-optimizer:-a-modern-optimizer-for-real-time-analytics-in-a-distributed-database.pdf) (2016) - Chen, Jack, et al.
### Transformation

- [Processing queries with quantifiers a horticultural approach](papers/transformation/processing-queries-with-quantifiers-a-horticultural-approach.pdf) (1983) - Dayal, Umeshwar.
- [Translating SQL into relational algebra: Optimization, semantics, and equivalence of SQL queries](papers/transformation/translating-sql-into-relational-algebra:-optimization,-semantics,-and-equivalence-of-sql-queries.pdf) (1985) - Ceri, Stefano, and Georg Gottlob.
- [Grammar-like Functional Rules for Representing Query Optimization Alternatives,](papers/transformation/grammar-like-functional-rules-for-representing-query-optimization-alternatives,.pdf) (1988) - Lohman, Guy M.
- [Query Optimization by Predicate Move-Around](papers/transformation/query-optimization-by-predicate-move-around.pdf) (1994) - Levy, Alon Y., Inderpal Singh Mumick, and Yehoshua Sagiv. .
- [Eager Aggregation and Lazy Aggregation](papers/transformation/eager-aggregation-and-lazy-aggregation.pdf) (1995) - Yan, Weipeng P., and Per-Bike Larson.
- [Parameterized Queries and Nesting Equivalences](papers/transformation/parameterized-queries-and-nesting-equivalences.pdf) (2000) - Galindo-Legaria, C. A.
- [Cost-based query transformation in Oracle](papers/transformation/cost-based-query-transformation-in-oracle.pdf) (2006) - Ahmed, Rafi, et al.
### Nested Query

- [Using semi-joins to solve relational queries](papers/nested-query/using-semi-joins-to-solve-relational-queries.pdf) (1981) - Bernstein, Philip A., and Dah-Ming W. Chiu.
- [On optimizing an SQL-like nested query](papers/nested-query/on-optimizing-an-sql-like-nested-query.pdf) (1982) - Kim, Won.
- [Optimization of nested queries in a distributed relational database](papers/nested-query/optimization-of-nested-queries-in-a-distributed-relational-database.pdf) (1984) - L&man, Guy M., et al.
- [SQL-like and Quel-like correlation queries with aggregates revisited](papers/nested-query/sql-like-and-quel-like-correlation-queries-with-aggregates-revisited.pdf) (1984) - Kiessling, Werner.
- [Translating SQL into relational algebra: Optimization, semantics, and equivalence of SQL queries](papers/nested-query/translating-sql-into-relational-algebra:-optimization,-semantics,-and-equivalence-of-sql-queries.pdf) (1985) - Ceri, Stefano, and Georg Gottlob.
- [Optimization of nested SQL queries revisited](papers/nested-query/optimization-of-nested-sql-queries-revisited.pdf) (1987) - Ganski, Richard A., and Harry KT Wong.
- [A Unitied Approach to Processing Queries That Contain Nested Subqueries, Aggregates, and Quantifiers](papers/nested-query/a-unitied-approach-to-processing-queries-that-contain-nested-subqueries,-aggregates,-and-quantifiers.pdf) (1987) - Dayal, Umeshwar.
- [Optimization of correlated SQL queries in a relational database management system](papers/nested-query/optimization-of-correlated-sql-queries-in-a-relational-database-management-system.pdf) (1998) - Jou, Michelle M., Ting Yu Leung, and Mir Hamid Pirahesh.
- [Orthogonal Optimization of Subqueries and Aggregation](papers/nested-query/orthogonal-optimization-of-subqueries-and-aggregation.pdf) (2001) - Galindo-Legaria, César, and Milind Joshi.
- [WinMagic : Subquery Elimination Using Window Aggregation](papers/nested-query/winmagic-:-subquery-elimination-using-window-aggregation.pdf) (2003) - Zuzarte, Calisto, et al. .
- [Execution strategies for SQL subqueries](papers/nested-query/execution-strategies-for-sql-subqueries.pdf) (2007) - Elhemali, Mostafa, et al.
- [Enhanced subquery optimizations in Oracle](papers/nested-query/enhanced-subquery-optimizations-in-oracle.pdf) (2009) - Bellamkonda, Srikanth, et al.
- [Unnesting Arbitrary Queries](papers/nested-query/unnesting-arbitrary-queries.pdf) (2015) - Neumann, Thomas, and Alfons Kemper.
### Functional Dependencies

- [Fundamental Techniques for Order Optimization](papers/functional-dependencies/fundamental-techniques-for-order-optimization.pdf) (1996) - Simmen, David, Eugene Shekita, and Timothy Malkemus. .
- [[Thesis] Exploiting Functional Dependence in Query Optimization](papers/functional-dependencies/[thesis]-exploiting-functional-dependence-in-query-optimization.pdf) (2000) - Paulley, Glenn Norman.
- [An Efficient Framework for Order Optimization](papers/functional-dependencies/an-efficient-framework-for-order-optimization.pdf) (2004) - Neumann, Thomas, and Guido Moerkotte.
- [Incorporating Partitioning and Parallel Plans into the SCOPE Optimizer](papers/functional-dependencies/incorporating-partitioning-and-parallel-plans-into-the-scope-optimizer.pdf) (2010) - Zhou, Jingren, Per-Ake Larson, and Ronnie Chaiken. .
- [Accelerating Queries with GroupBy and Join by Group join](papers/functional-dependencies/accelerating-queries-with-groupby-and-join-by-group-join.pdf) (2011) - Moerkotte, Guido, and Thomas Neumann.
### Join Order

- [Access paths in the" Abe" statistical query facility](papers/join-order/access-paths-in-the"-abe"-statistical-query-facility.pdf) (1982) - Klug, Anthony.
- [Extending the Algebraic Framework of Query Processing to Handle Outerjoins](papers/join-order/extending-the-algebraic-framework-of-query-processing-to-handle-outerjoins.pdf) (1984) - RosenthaI, A., and D. Reiner.
- [Analysis of Two Existing and One New Dynamic Programming Algorithm for the Generation of Optimal Bushy Join Trees without Cross Products](papers/join-order/analysis-of-two-existing-and-one-new-dynamic-programming-algorithm-for-the-generation-of-optimal-bushy-join-trees-without-cross-products.pdf) (2006) - Moerkotte, Guido, and Thomas Neumann.
- [Dynamic programming strikes back](papers/join-order/dynamic-programming-strikes-back.pdf) (2008) - Moerkotte, Guido, and Thomas Neumann.
- [On the Correct and Complete Enumeration of the Core Search Space](papers/join-order/on-the-correct-and-complete-enumeration-of-the-core-search-space.pdf) (2013) - Moerkotte, Guido, Pit Fender, and Marius Eich.
- [How Good Are Query Optimizers, Really?](papers/join-order/how-good-are-query-optimizers,-really?.pdf) (2015) - Leis, Viktor, et al.
- [The Complete Story of Joins](papers/join-order/the-complete-story-of-joins.pdf) (2017) - Neumann, Thomas, Viktor Leis, and Alfons Kemper.
- [Improving Join Reorderability with Compensation Operators](papers/join-order/improving-join-reorderability-with-compensation-operators.pdf) (2018) - Wang, TaiNing, and Chee-Yong Chan.
- [Adaptive Optimization of Very Large Join Queries](papers/join-order/adaptive-optimization-of-very-large-join-queries.pdf) (2018) - Neumann, Thomas, and Bernhard Radke.
### Cost Model

- [Modelling Costs for a MM-DBMS](papers/cost-model/modelling-costs-for-a-mm-dbms.pdf) (1996) - Listgarten, Sherry, and Marie-Anne Neimat.
- [SEEKing the truth about ad hoc join costs](papers/cost-model/seeking-the-truth-about-ad-hoc-join-costs.pdf) (1997) - Haas, Laura M., et al.
- [Approximation Schemes for Many-Objective Query Optimization](papers/cost-model/approximation-schemes-for-many-objective-query-optimization.pdf) (2014) - Trummer, Immanuel, and Christoph Koch.
- [Multi-Objective Parametric Query Optimization](papers/cost-model/multi-objective-parametric-query-optimization.pdf) (2015) - Trummer, Immanuel, and Christoph Koch.
### Statistics

- [Accurate Estimation of the Number of Tuples Satisfying a Condition](papers/statistics/accurate-estimation-of-the-number-of-tuples-satisfying-a-condition.pdf) (1984) - Piatetsky-Shapiro, Gregory, and Charles Connell.
- [Optimal Histograms for Limiting Worst-Case Error Propagation in the Size of Join Results](papers/statistics/optimal-histograms-for-limiting-worst-case-error-propagation-in-the-size-of-join-results.pdf) (1993) - Ioannidis, Yannis E., and Stavros Christodoulakis.
- [Universality of Serial Histograms](papers/statistics/universality-of-serial-histograms.pdf) (1993) - Ioannidis, Yannis E.
- [Balancing Histogram Optimality and Practicality for Query Result Size Estimation](papers/statistics/balancing-histogram-optimality-and-practicality-for-query-result-size-estimation.pdf) (1995) - Ioannidis, Yannis E., and Viswanath Poosala.
- [Improved Histograms for Selectivity Estimation of Range Predicates](papers/statistics/improved-histograms-for-selectivity-estimation-of-range-predicates.pdf) (1996) - Poosala, Viswanath, et al.
- [The History of Histograms](papers/statistics/the-history-of-histograms.pdf) (2003) - Ioannidis, Yannis.
- [Automated Statistics Collection in DB2 UDB](papers/statistics/automated-statistics-collection-in-db2-udb.pdf) (2004) - Aboulnaga, Ashraf, et al.
- [Adaptive Query Processing in the Looking Glass](papers/statistics/adaptive-query-processing-in-the-looking-glass.pdf) (2005) - Babu, Shivnath, and Pedro Bizarro.
- [Optimizer plan change management: improved stability and performance in Oracle 11g](papers/statistics/optimizer-plan-change-management:-improved-stability-and-performance-in-oracle-11g.pdf) (2008) - Ziauddin, Mohamed, et al.
- [Histograms Reloaded: The Merits of Bucket Diversity](papers/statistics/histograms-reloaded:-the-merits-of-bucket-diversity.pdf) (2010) - Kanne, Carl-Christian, and Guido Moerkotte.
- [Synopses for Massive Data: Samples, Histograms, Wavelets, Sketches](papers/statistics/synopses-for-massive-data:-samples,-histograms,-wavelets,-sketches.pdf) (2011) - Cormode, Graham, et al.
- [Exploiting Ordered Dictionaries to Efficiently Construct Histograms with Q-Error Guarantees in SAP HANA](papers/statistics/exploiting-ordered-dictionaries-to-efficiently-construct-histograms-with-q-error-guarantees-in-sap-hana.pdf) (2014) - Moerkotte, Guido, et al.
- [Adaptive Statistics in Oracle 12c](papers/statistics/adaptive-statistics-in-oracle-12c.pdf) (2017) - Chakkappen, Sunil, et al.
### Probabilistic Counting

- [Towards Estimation Error Guarantees for Distinct Values](papers/probabilistic-counting/towards-estimation-error-guarantees-for-distinct-values.pdf) (2000) - Charikar, Moses, et al.
- [Distinct Sampling for Highly-Accurate Answers to Distinct Values Queries and Event Reports](papers/probabilistic-counting/distinct-sampling-for-highly-accurate-answers-to-distinct-values-queries-and-event-reports.pdf) (2001) - Gibbons, Phillip B.
- [An Improved Data Stream Summary: The Count-Min Sketch and its Applications, Journal of Algorithms](papers/probabilistic-counting/an-improved-data-stream-summary:-the-count-min-sketch-and-its-applications,-journal-of-algorithms.pdf) (2005) - Cormode, Graham, and Shan Muthukrishnan.
- [New Estimation Algorithms for Streaming Data: Count-min Can Do More](papers/probabilistic-counting/new-estimation-algorithms-for-streaming-data:-count-min-can-do-more.pdf) (2007) - Deng, Fan, and Davood Rafiei.
- [Preventing Bad Plans by Bounding the Impact of Cardinality Estimation Errors](papers/probabilistic-counting/preventing-bad-plans-by-bounding-the-impact-of-cardinality-estimation-errors.pdf) (2009) - Moerkotte, Guido, Thomas Neumann, and Gabriele Steidl.
- [Pessimistic Cardinality Estimation: Tighter Upper Bounds for Intermediate Join Cardinalities](papers/probabilistic-counting/pessimistic-cardinality-estimation:-tighter-upper-bounds-for-intermediate-join-cardinalities.pdf) (2019) - Cai, Walter, Magdalena Balazinska, and Dan Suciu.
- [Deep Unsupervised Cardinality Estimation](papers/probabilistic-counting/deep-unsupervised-cardinality-estimation.pdf) (2019) - Yang, Zongheng, et al.
- [NeuroCard: One Cardinality Estimator for All Tables](papers/probabilistic-counting/neurocard:-one-cardinality-estimator-for-all-tables.pdf) (2020) - Yang, Zongheng, et al.
### Execution Engine

- [QueryEvaluationTechniquesfor LargeDatabas](papers/execution-engine/queryevaluationtechniquesfor-largedatabas.pdf) (1993) - Graefe G.
- [Volcano - An Extensible and Parallel Query Evaluation System](papers/execution-engine/volcano---an-extensible-and-parallel-query-evaluation-system.pdf) (1994) - Graefe, Goetz.
- [MonetDB/X100: Hyper-Pipelining Query Execution](papers/execution-engine/monetdb-x100:-hyper-pipelining-query-execution.pdf) (2005) - Boncz, Peter A., Marcin Zukowski, and Niels Nes.
- [Efficiently Compiling Efficient Query Plans for Modern Hardware](papers/execution-engine/efficiently-compiling-efficient-query-plans-for-modern-hardware.pdf) (2011) - Neumann, Thomas.
- [Multi-Core, Main-Memory Joins: Sort vs. Hash Revisited](papers/execution-engine/multi-core,-main-memory-joins:-sort-vs.-hash-revisited.pdf) (2013) - Balkesen, Cagri, et al.
- [Morsel-Driven Parallelism: A NUMA-Aware Query Evaluation Framework for the Many-Core Age](papers/execution-engine/morsel-driven-parallelism:-a-numa-aware-query-evaluation-framework-for-the-many-core-age.pdf) (2014) - Leis, Viktor, et al.
- [Relaxed Operator Fusion for In-Memory Databases: Making Compilation, Vectorization, and Prefetching Work Together At Last](papers/execution-engine/relaxed-operator-fusion-for-in-memory-databases:-making-compilation,-vectorization,-and-prefetching-work-together-at-last.pdf) (2017) - Menon, Prashanth, Todd C. Mowry, and Andrew Pavlo.
- [Looking Ahead Makes Query Plans Robust](papers/execution-engine/looking-ahead-makes-query-plans-robust.pdf) (2017) - Zhu, Jianqiao, et al.
- [Everything You Always Wanted to Know About Compiled and Vectorized Queries But Were Afraid to Ask](papers/execution-engine/everything-you-always-wanted-to-know-about-compiled-and-vectorized-queries-but-were-afraid-to-ask.pdf) (2018) - Kersten, Timo, et al.
- [SuRF: Practical Range Query Filtering with Fast Succinct Tries](papers/execution-engine/surf:-practical-range-query-filtering-with-fast-succinct-tries.pdf) (2018) - Zhang, Huanchen, et al.
- [Adaptive Execution of Compiled Queries](papers/execution-engine/adaptive-execution-of-compiled-queries.pdf) (2018) - Kohn, André, Viktor Leis, and Thomas Neumann.
### MPP Optimizations

- [DB2 Parallel Edition](papers/mpp-optimizations/db2-parallel-edition.pdf) (1995) - Baru, Chaitanya K., et al.
- [Parallel SQL execution in Oracle 10g](papers/mpp-optimizations/parallel-sql-execution-in-oracle-10g.pdf) (2004) - Cruanes, Thierry, Benoit Dageville, and Bhaskar Ghosh.
- [Query Optimization in Microsoft SQL Server PDW](papers/mpp-optimizations/query-optimization-in-microsoft-sql-server-pdw.pdf) (2012) - Shankar, Srinath, et al.
- [Adaptive and big data scale parallel execution in Oracle](papers/mpp-optimizations/adaptive-and-big-data-scale-parallel-execution-in-oracle.pdf) (2013) - Bellamkonda, Srikanth, et al.
- [Optimizing Queries over Partitioned Tables in MPP Systems](papers/mpp-optimizations/optimizing-queries-over-partitioned-tables-in-mpp-systems.pdf) (2014) - Antova, Lyublena, et al.
## Storage Engine

### Storage Structure

- [The Ubiquitous B-Tree](papers/storage-structure/the-ubiquitous-b-tree.pdf) (1979) - Comer, Douglas.
- [The 5 Minute Rule for Trading Memory for Disc Accesses and the 5 Byte Rule for Trading Memory for CPU Time](papers/storage-structure/the-5-minute-rule-for-trading-memory-for-disc-accesses-and-the-5-byte-rule-for-trading-memory-for-cpu-time.pdf) (1987) - Gray, Jim, and Franco Putzolu.
- [The Log-Structured Merge-Tree (LSM-Tree)](papers/storage-structure/the-log-structured-merge-tree-(lsm-tree).pdf) (1996) - O’Neil, Patrick, et al.
- [The five-minute rule ten years later, and other computer storage rules of thumb](papers/storage-structure/the-five-minute-rule-ten-years-later,-and-other-computer-storage-rules-of-thumb.pdf) (1997) - Gray, Jim, and Goetz Graefe.
- [The Five Minute Rule 20 Years Later and How Flash Memory Changes the Rules](papers/storage-structure/the-five-minute-rule-20-years-later-and-how-flash-memory-changes-the-rules.pdf) (2008) - Graefe, Goetz.
- [A Comparison of Fractal Trees to Log-Structured Merge (LSM) Trees](papers/storage-structure/a-comparison-of-fractal-trees-to-log-structured-merge-(lsm)-trees.pdf) (2014) - Kuszmaul, Bradley C.
- [Design Tradeoffs of Data Access Methods](papers/storage-structure/design-tradeoffs-of-data-access-methods.pdf) (2016) - Athanassoulis, Manos, and Stratos Idreos. .
- [Designing Access Methods: The RUM Conjecture](papers/storage-structure/designing-access-methods:-the-rum-conjecture.pdf) (2016) - Athanassoulis, Manos, et al.
- [The five minute rule thirty years later and its impact on the storage hierarchy](papers/storage-structure/the-five-minute-rule-thirty-years-later-and-its-impact-on-the-storage-hierarchy.pdf) (2017) - Appuswamy, Raja, et al.
- [WiscKey: Separating Keys from Values in SSD-conscious Storage](papers/storage-structure/wisckey:-separating-keys-from-values-in-ssd-conscious-storage.pdf) (2017) - Lu, Lanyue, et al.
- [Managing Non-Volatile Memory in Database Systems](papers/storage-structure/managing-non-volatile-memory-in-database-systems.pdf) (2018) - van Renen, Alexander, et al.
- [LeanStore: In-Memory Data Management Beyond Main Memory](papers/storage-structure/leanstore:-in-memory-data-management-beyond-main-memory.pdf) (2018) - Leis, Viktor, et al.
- [The Case for Learned Index Structures](papers/storage-structure/the-case-for-learned-index-structures.pdf) (2018) - Kraska, Tim, et al.
- [LSM-based Storage Techniques: A Survey](papers/storage-structure/lsm-based-storage-techniques:-a-survey.pdf) (2019) - Luo, Chen, and Michael J. Carey.
- [Learning Multi-dimensional Indexes](papers/storage-structure/learning-multi-dimensional-indexes.pdf) (2019) - Nathan, Vikram, et al.
- [Umbra: A Disk-Based System with In-Memory Performance](papers/storage-structure/umbra:-a-disk-based-system-with-in-memory-performance.pdf) (2020) - Neumann, Thomas, and Michael J. Freitag.
- [XIndex: A Scalable Learned Index for Multicore Data Storage](papers/storage-structure/xindex:-a-scalable-learned-index-for-multicore-data-storage.pdf) (2020) - Tang, Chuzhe, et al.
- [The PGM-index: a fully-dynamic compressed learned index with provable worst-case bounds](papers/storage-structure/the-pgm-index:-a-fully-dynamic-compressed-learned-index-with-provable-worst-case-bounds.pdf) (2020) - Ferragina, Paolo, and Giorgio Vinciguerra.
- [From WiscKey to Bourbon: A Learned Index for Log-Structured Merge Trees](papers/storage-structure/from-wisckey-to-bourbon:-a-learned-index-for-log-structured-merge-trees.pdf) (2020) - Dai, Yifan, et al. .
### Transaction

- [The Notions of Consistency and Predicate Locks in a Database System](papers/transaction/the-notions-of-consistency-and-predicate-locks-in-a-database-system.pdf) (1976) - Eswaran, Kapali P., et al.
- [Concurrency Control in Distributed Database Systems](papers/transaction/concurrency-control-in-distributed-database-systems.pdf) (1981) - Bernstein, Philip A., and Nathan Goodman.
- [On Optimistic Methods for Concurrency Control](papers/transaction/on-optimistic-methods-for-concurrency-control.pdf) (1981) - Kung, Hsiang-Tsung, and John T. Robinson.
- [Principles of transaction-oriented database recovery](papers/transaction/principles-of-transaction-oriented-database-recovery.pdf) (1983) - Haerder, Theo, and Andreas Reuter.
- [Multiversion Concurrency Control - Theory and Algorithms](papers/transaction/multiversion-concurrency-control---theory-and-algorithms.pdf) (1983) - Bernstein, Philip A., and Nathan Goodman.
- [ARIES: A transaction recovery method supporting fine-granularity locking and partial rollbacks using write-ahead logging](papers/transaction/aries:-a-transaction-recovery-method-supporting-fine-granularity-locking-and-partial-rollbacks-using-write-ahead-logging.pdf) (1992) - Mohan C, Haderle D, Lindsay B, et al.
- [A Critique of ANSI SQL Isolation Levels](papers/transaction/a-critique-of-ansi-sql-isolation-levels.pdf) (1995) - Berenson, Hal, et al.
- [Generalized Isolation Level Definitions](papers/transaction/generalized-isolation-level-definitions.pdf) (2000) - Adya, Atul, Barbara Liskov, and Patrick O'Neil.
- [Serializable Snapshot Isolation in PostgreSQL](papers/transaction/serializable-snapshot-isolation-in-postgresql.pdf) (2012) - Ports, Dan RK, and Kevin Grittner.
- [Calvin: Fast Distributed Transactions for Partitioned Database Systems](papers/transaction/calvin:-fast-distributed-transactions-for-partitioned-database-systems.pdf) (2012) - Thomson, Alexander, et al.
- [MaaT: effective and scalable coordination of distributed transactions in the cloud](papers/transaction/maat:-effective-and-scalable-coordination-of-distributed-transactions-in-the-cloud.pdf) (2014) - Mahmoud, Hatem A., et al.
- [Staring into the Abyss: An Evaluation of Concurrency Control with One Thousand Cores](papers/transaction/staring-into-the-abyss:-an-evaluation-of-concurrency-control-with-one-thousand-cores.pdf) (2014) - Yu, Xiangyao, et al.
- [An Evaluation of the Advantages and Disadvantages of Deterministic Database Systems](papers/transaction/an-evaluation-of-the-advantages-and-disadvantages-of-deterministic-database-systems.pdf) (2014) - Ren, Kun, Alexander Thomson, and Daniel J. Abadi.
- [Fast Serializable Multi-Version Concurrency Control for Main-Memory Database Systems](papers/transaction/fast-serializable-multi-version-concurrency-control-for-main-memory-database-systems.pdf) (2015) - Neumann, Thomas, Tobias Mühlbauer, and Alfons Kemper.
- [An Empirical Evaluation of In-Memory Multi-Version Concurrency Control](papers/transaction/an-empirical-evaluation-of-in-memory-multi-version-concurrency-control.pdf) (2017) - Wu, Yingjun, et al.
- [An Evaluation of Distributed Concurrency Control](papers/transaction/an-evaluation-of-distributed-concurrency-control.pdf) (2017) - Harding, Rachael, et al.
- [Scalable Garbage Collection for In-Memory MVCC Systems](papers/transaction/scalable-garbage-collection-for-in-memory-mvcc-systems.pdf) (2019) - Böttcher, Jan, et al.
### Scheduling

- [Automated Demand-driven Resource Scaling in Relational Database-as-a-Service](papers/scheduling/automated-demand-driven-resource-scaling-in-relational-database-as-a-service.pdf) (2016) - Das, Sudipto, et al.
- [Autoscaling Tiered Cloud Storage in Anna](papers/scheduling/autoscaling-tiered-cloud-storage-in-anna.pdf) (2019) - Wu, Chenggang, Vikram Sreekanti, and Joseph M. Hellerstein.
- [Adaptive HTAP through Elastic Resource Scheduling](papers/scheduling/adaptive-htap-through-elastic-resource-scheduling.pdf) (2020) - Raza, Aunn, et al.
- [MorphoSys: Automatic Physical Design Metamorphosis for Distributed Database Systems](papers/scheduling/morphosys:-automatic-physical-design-metamorphosis-for-distributed-database-systems.pdf) (2020) - Abebe, Michael, Brad Glasbergen, and Khuzaima Daudjee.
