[![README autogen](https://github.com/lonng/db-papers/actions/workflows/autogen.yml/badge.svg)](https://github.com/lonng/db-papers/actions/workflows/autogen.yml)

---

# Database Papers

This is a comprehensive list of papers on database theory for understanding and building database systems. It covers various aspects of database systems, including the essential theoretical background, classic system design, and multiple modules within the database.

The list is organized into different categories and subcategories for easy navigation. Each paper is accompanied by a title, author, and publication year, along with a link to the full text if available.

This collection serves as a learning and training resource primarily for the Tencent Cloud Database Team and is also open to external researchers, students, and learners interested in database systems.

**In case you are reading this and making the effort to comprehend these papers, we would really like to have a conversation with you regarding opportunities at Tencent Cloud Database Team ([@Henry L.](https://x.com/henrylonng)).**

## Contribution

This list is generated from a Sheet document automatically. If you have any suggestions or would like to contribute to this list, please feel free to file an issue. And we will update our sheet to make the chagnes available for public.

Any contribution that can help improve this list and make it more comprehensive and useful to the community are welcome. Here are some ways you can contribute:

- **Add a new paper**: If you have a paper that you think should be included in this list, please file an issue to provide the paper's title, author, publication year, and a link to the full text (if available).
- **Update an existing paper**: If you find any errors or outdated information in the list, please file an issue to provide the correct information.
- **Remove a paper**: If you think a paper is no longer relevant or useful, please file an issue to suggest its removal.
- **General suggestions**: If you have any general suggestions or feedback on how to improve this list, please file an issue to share your thoughts.

## Table of Contents

- [Database Papers](#database-papers)
  - [Contribution](#contribution)
  - [Table of Contents](#table-of-contents)
  - [Basics](#basics)
    - [Essentials](#essentials)
    - [Consensus](#consensus)
    - [Consistency](#consistency)
  - [System Design](#system-design)
    - [Architecture](#architecture)
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
    - [Parallel Execution](#parallel-execution)
  - [Storage Engine](#storage-engine)
    - [Storage Media](#storage-media)
    - [Storage Structure](#storage-structure)
    - [Transaction](#transaction)
    - [Scheduling](#scheduling)
  - [Miscellaneous](#miscellaneous)
    - [Workload](#workload)
    - [Network](#network)
    - [Quality](#quality)
    - [Diagnosis and Tuning](#diagnosis-and-tuning)

## Basics

### Essentials

- [A Relational Model Of Data For Large Shared Data Banks](papers/essentials/a-relational-model-of-data-for-large-shared-data-banks.pdf) (1970) - Codd, Edgar F.
- [Sequel: A Structured English Query Language](papers/essentials/sequel:-a-structured-english-query-language.pdf) (1974) - Chamberlin, Donald D., and Raymond F. Boyce.
- [Ingres: A Relational Data Base System](papers/essentials/ingres:-a-relational-data-base-system.pdf) (1975) - Held, G. D., M. R. Stonebraker, and Eugene Wong.
- [Extending The Database Relational Model To Capture More Meaning](papers/essentials/extending-the-database-relational-model-to-capture-more-meaning.pdf) (1979) - Codd, Edgar F.
- [A Critique Of The Sql Database Language](papers/essentials/a-critique-of-the-sql-database-language.pdf) (1984) - Date, C. J.
- [A Critique Of Snapshot Isolation](https://dl.acm.org/doi/pdf/10.1145/2168836.2168853) (2012) - Yabandeh M, Gómez Ferro D.

### Consensus

- [The Part-Time Parliament](https://dl.acm.org/doi/pdf/10.1145/3335772.3335939) (1998) - Lamport, Leslie.
- [Paxos Made Simple](https://www.microsoft.com/en-us/research/publication/2016/12/paxos-simple-Copy.pdf) (2001) - Lamport, Leslie.
- [Consensus: Bridging Theory And Practice](papers/consensus/consensus:-bridging-theory-and-practice.pdf) (2014) - Ongaro, Diego.
- [In Search Of An Understandable Consensus Algorithm (Extended Version)](papers/consensus/in-search-of-an-understandable-consensus-algorithm-(extended-version).pdf) (2014) - Ongaro, Diego, and John Ousterhout.
- [Distributed Consensus Revised](papers/consensus/distributed-consensus-revised.pdf) (2019) - Howard, Heidi.
- [A Generalised Solution To Distributed Consensus](papers/consensus/a-generalised-solution-to-distributed-consensus.pdf) (2019) - Howard, Heidi, and Richard Mortier.
- [Paxos Vs Raft: Have We Reached Consensus On Distributed Consensus?](https://dl.acm.org/doi/pdf/10.1145/3380787.3393681) (2020) - Howard, Heidi, and Richard Mortier.

### Consistency

- [Consistency Tradeoffs In Modern Distributed Database System Design](papers/consistency/consistency-tradeoffs-in-modern-distributed-database-system-design.pdf) (2012) - Abadi, Daniel.
- [Logical Physical Clocks And Consistent Snapshots In Globally Distributed Databases](papers/consistency/logical-physical-clocks-and-consistent-snapshots-in-globally-distributed-databases.pdf) (2014) - Kulkarni S S, Demirbas M, Madappa D, et al.
- [Ark: A Real-World Consensus Implementation](papers/consistency/ark:-a-real-world-consensus-implementation.pdf) (2014) - Kasheff, Zardosht, and Leif Walsh.
- [Polarfs: An Ultra-Low Latency And Failure Resilient Distributed File System For Shared Storage Cloud Database](https://dl.acm.org/doi/pdf/10.14778/3229863.3229872) (2018) - Cao, Wei, et al.
- [Anna: A Kvs For Any Scale](papers/consistency/anna:-a-kvs-for-any-scale.pdf) (2018) - Wu, Chenggang, et al.
- [Strong And Efficient Consistency With Consistency-Aware Durability](https://dl.acm.org/doi/pdf/10.1145/3423138) (2021) - Ganesan, Aishwarya, et al.

## System Design

### Architecture

- [Architecture Of A Database System. Foundations And Trends In Databases](papers/architecture/architecture-of-a-database-system.-foundations-and-trends-in-databases.pdf) (2007) - Hellerstein J M, Stonebraker M, Hamilton J.

### RDBMS

- [System R: Relational Approach To Database Management](https://dl.acm.org/doi/pdf/10.1145/320455.320457) (1976) - Astrahan, Morton M., et al.
- [The Design And Implementation Of Ingres](https://dl.acm.org/doi/10.1145/320473.320476) (1976) - Stonebraker, Michael, et al.
- [The Design Of Postgres](https://dl.acm.org/doi/pdf/10.1145/16856.16888) (1986) - Stonebraker, Michael, and Lawrence A. Rowe.
- [Query Processing In Main Memory Database Management Systems](https://dl.acm.org/doi/pdf/10.1145/16894.16878) (1986) - Lehman, Tobin J., and Michael J. Carey.
- [Megastore: Providing Scalable, Highly Available Storage For Interactive Services](papers/rdbms/megastore:-providing-scalable,-highly-available-storage-for-interactive-services.pdf) (2011) - Baker J, Bond C, Corbett J C, et al.
- [Spanner: Google's Globally Distributed Database](https://dl.acm.org/doi/pdf/10.1145/2491245) (2013) - Corbett, James C., et al.
- [Online, Asynchronous Schema Change In F1](https://dl.acm.org/doi/pdf/10.14778/2536222.2536230) (2013) - Rae, Ian, et al.
- [Amazon Aurora: Design Considerations For High Throughput Cloud-Native Relational Databases](https://dl.acm.org/doi/pdf/10.1145/3035918.3056101) (2017) - Verbitski, Alexandre, et al.
- [Looking Back At Postgres](papers/rdbms/looking-back-at-postgres.pdf) (2019) - Hellerstein, Joseph M.
- [Cockroachdb: The Resilient Geo-Distributed Sql Database](https://dl.acm.org/doi/pdf/10.1145/3318464.3386134) (2020) - Taft, Rebecca, et al.
- [F1 Lightning: Htap As A Service](https://dl.acm.org/doi/pdf/10.14778/3415478.3415553) (2020) - Yang, Jiacheng, et al.
- [Tidb: A Raft-Based Htap Database](https://dl.acm.org/doi/pdf/10.14778/3415478.3415535) (2020) - Huang, Dongxu, et al.
- [Polardb Serverless: A Cloud Native Database For Disaggregated Data Centers](https://dl.acm.org/doi/pdf/10.1145/3448016.3457560) (2021) - Cao, Wei, et al.

### NoSQL

- [Bigtable: A Distributed Storage System For Structured Data](https://dl.acm.org/doi/pdf/10.1145/1365815.1365816) (2006) - Chang, Fay, et al.
- [Dynamo: Amazon’s Highly Available Key-Value Store](https://dl.acm.org/doi/pdf/10.1145/1323293.1294281) (2007) - DeCandia, Giuseppe, et al.
- [Pnuts: Yahoo!’S Hosted Data Serving Platform](https://dl.acm.org/doi/pdf/10.14778/1454159.1454167) (2008) - Cooper, Brian F., et al.
- [Cassandra - A Decentralized Structured Storage System](https://dl.acm.org/doi/pdf/10.1145/1773912.1773922) (2010) - Lakshman, Avinash, and Prashant Malik.
- [Windows Azure Storage: A Highly Available Cloud Storage Service With Strong Consistency](https://dl.acm.org/doi/pdf/10.1145/2043556.2043571) (2011) - Calder, Brad, et al.
- [Azure Data Lake Store: A Hyperscale Distributed File Service For Big Data Analytics](https://dl.acm.org/doi/pdf/10.1145/3035918.3056100) (2017) - Ramakrishnan, Raghu, et al.
- [Pnuts To Sherpa: Lessons From Yahoo!’S Cloud Database](https://dl.acm.org/doi/pdf/10.14778/3352063.3352146) (2019) - Cooper, Brian F., et al.

## SQL Engine

### Optimizer Framework

- [Access Path Selection In A Relational Database Management System](https://dl.acm.org/doi/pdf/10.1145/582095.582099) (1979) - Selinger, P. Griffiths, et al.
- [Query Optimization By Simulated Annealing](https://dl.acm.org/doi/pdf/10.1145/38713.38722) (1987) - Ioannidis, Yannis E., and Eugene Wong.
- [The Exodus Optimizer Generator](https://dl.acm.org/doi/pdf/10.1145/38713.38734) (1987) - Graefe, Goetz, and David J. DeWitt.
- [Extensible/Rule Based Query Rewrite Optimization In Starburst](https://dl.acm.org/doi/pdf/10.1145/141484.130294) (1992) - Pirahesh, Hamid, Joseph M. Hellerstein, and Waqar Hasan.
- [The Volcano Optimizer Generator- Extensibility And Efficient Search](papers/optimizer-framework/the-volcano-optimizer-generator--extensibility-and-efficient-search.pdf) (1993) - Graefe, Goetz, and William J. McKenna.
- [The Cascades Framework For Query Optimization](papers/optimizer-framework/the-cascades-framework-for-query-optimization.pdf) (1995) - Graefe, Goetz.
- [An Overview Of Query Optimization In Relational Systems](https://dl.acm.org/doi/pdf/10.1145/275487.275492) (1998) - Chaudhuri, Surajit.
- [Robust Query Processing Through Progressive Optimization](https://dl.acm.org/doi/pdf/10.1145/1007568.1007642) (2004) - Markl, Volker, et al.
- [Orca: A Modular Query Optimizer Architecture For Big Data](https://dl.acm.org/doi/pdf/10.1145/2588555.2595637) (2014) - Soliman, Mohamed A., et al.
- [Parallelizing Query Optimization On Shared-Nothing Architectures](papers/optimizer-framework/parallelizing-query-optimization-on-shared-nothing-architectures.pdf) (2015) - Trummer, Immanuel, and Christoph Koch.
- [The Memsql Query Optimizer: A Modern Optimizer For Real-Time Analytics In A Distributed Database](https://dl.acm.org/doi/pdf/10.14778/3007263.3007277) (2016) - Chen, Jack, et al.

### Transformation

- [Processing Queries With Quantifiers A Horticultural Approach](https://dl.acm.org/doi/pdf/10.1145/588058.588075) (1983) - Dayal, Umeshwar.
- [Translating Sql Into Relational Algebra: Optimization, Semantics, And Equivalence Of Sql Queries](https://www.academia.edu/download/50687636/tse.1985.23222320161202-29901-8u86ef.pdf) (1985) - Ceri, Stefano, and Georg Gottlob.
- [Grammar-Like Functional Rules For Representing Query Optimization Alternatives,](https://dl.acm.org/doi/pdf/10.1145/971701.50204) (1988) - Lohman, Guy M.
- [Query Optimization By Predicate Move-Around](https://www.researchgate.net/profile/Inderpal-Mumick/publication/2754592_Query_Optimization_by_Predicate_Move-Around/links/0f317534d437e49755000000/Query-Optimization-by-Predicate-Move-Around.pdf) (1994) - Levy, Alon Y., Inderpal Singh Mumick, and Yehoshua Sagiv.
- [Eager Aggregation And Lazy Aggregation](https://www.researchgate.net/profile/Per-Ake-Larson/publication/2733082_Eager_Aggregation_and_Lazy_Aggregation/links/02bfe50ce6de3dad7c000000/Eager-Aggregation-and-Lazy-Aggregation.pdf) (1995) - Yan, Weipeng P., and Per-Bike Larson.
- [Parameterized Queries And Nesting Equivalences](https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/tr-2000-31.pdf) (2000) - Galindo-Legaria, C. A.
- [Cost-Based Query Transformation In Oracle](https://www.researchgate.net/profile/Rafi-Ahmed-2/publication/221311318_Cost-Based_Query_Transformation_in_Oracle/links/572bbc5e08aef7c7e2c6b829/Cost-Based-Query-Transformation-in-Oracle.pdf) (2006) - Ahmed, Rafi, et al.

### Nested Query

- [Using Semi-Joins To Solve Relational Queries](https://dl.acm.org/doi/pdf/10.1145/322234.322238) (1981) - Bernstein, Philip A., and Dah-Ming W. Chiu.
- [On Optimizing An Sql-Like Nested Query](https://dl.acm.org/doi/pdf/10.1145/319732.319745) (1982) - Kim, Won.
- [Optimization Of Nested Queries In A Distributed Relational Database](papers/nested-query/optimization-of-nested-queries-in-a-distributed-relational-database.pdf) (1984) - L&man, Guy M., et al.
- [Sql-Like And Quel-Like Correlation Queries With Aggregates Revisited](papers/nested-query/sql-like-and-quel-like-correlation-queries-with-aggregates-revisited.pdf) (1984) - Kiessling, Werner.
- [Translating Sql Into Relational Algebra: Optimization, Semantics, And Equivalence Of Sql Queries](https://www.academia.edu/download/50687636/tse.1985.23222320161202-29901-8u86ef.pdf) (1985) - Ceri, Stefano, and Georg Gottlob.
- [Optimization Of Nested Sql Queries Revisited](https://dl.acm.org/doi/pdf/10.1145/38714.38723) (1987) - Ganski, Richard A., and Harry KT Wong.
- [A Unitied Approach To Processing Queries That Contain Nested Subqueries, Aggregates, And Quantifiers](papers/nested-query/a-unitied-approach-to-processing-queries-that-contain-nested-subqueries,-aggregates,-and-quantifiers.pdf) (1987) - Dayal, Umeshwar.
- [Orthogonal Optimization Of Subqueries And Aggregation](https://dl.acm.org/doi/pdf/10.1145/376284.375748) (2001) - Galindo-Legaria, César, and Milind Joshi.
- [Winmagic : Subquery Elimination Using Window Aggregation](https://dl.acm.org/doi/pdf/10.1145/872757.872840) (2003) - Zuzarte, Calisto, et al.
- [Execution Strategies For Sql Subqueries](https://dl.acm.org/doi/pdf/10.1145/1247480.1247598) (2007) - Elhemali, Mostafa, et al.
- [Enhanced Subquery Optimizations In Oracle](https://dl.acm.org/doi/pdf/10.14778/1687553.1687563) (2009) - Bellamkonda, Srikanth, et al.
- [Unnesting Arbitrary Queries](papers/nested-query/unnesting-arbitrary-queries.pdf) (2015) - Neumann, Thomas, and Alfons Kemper.
- [The Complete Story Of Joins](papers/nested-query/the-complete-story-of-joins.pdf) (2017) - Neumann, Thomas, Viktor Leis, and Alfons Kemper.

### Functional Dependencies

- [Fundamental Techniques For Order Optimization](https://dl.acm.org/doi/pdf/10.1145/233269.233320) (1996) - Simmen, David, Eugene Shekita, and Timothy Malkemus.
- [[Thesis] Exploiting Functional Dependence In Query Optimization](papers/functional-dependencies/[thesis]-exploiting-functional-dependence-in-query-optimization.pdf) (2000) - Paulley, Glenn Norman.
- [An Efficient Framework For Order Optimization](papers/functional-dependencies/an-efficient-framework-for-order-optimization.pdf) (2004) - Neumann, Thomas, and Guido Moerkotte.
- [Incorporating Partitioning And Parallel Plans Into The Scope Optimizer](papers/functional-dependencies/incorporating-partitioning-and-parallel-plans-into-the-scope-optimizer.pdf) (2010) - Zhou, Jingren, Per-Ake Larson, and Ronnie Chaiken.
- [Accelerating Queries With Groupby And Join By Group Join](https://dl.acm.org/doi/pdf/10.14778/3402707.3402723) (2011) - Moerkotte, Guido, and Thomas Neumann.

### Join Order

- [Access Paths In The" Abe" Statistical Query Facility](https://dl.acm.org/doi/pdf/10.1145/582353.582382) (1982) - Klug, Anthony.
- [Extending The Algebraic Framework Of Query Processing To Handle Outerjoins](papers/join-order/extending-the-algebraic-framework-of-query-processing-to-handle-outerjoins.pdf) (1984) - RosenthaI, A., and D. Reiner.
- [Outerjoin Simplication And Reordering For Query Optimization](https://dl.acm.org/doi/pdf/10.1145/244810.244812) (1993) - Galindo-Legaria C, Rosenthal A.
- [Hypergraph Based Reorderings Of Outer Join Queries With Complex Predicates](https://dl.acm.org/doi/pdf/10.1145/223784.223847) (1995) - Bhargava G, Goel P, Iyer B.
- [Rapid Bushy Join-Order Optimization With Cartesian Products.](https://dl.acm.org/doi/pdf/10.1145/235968.233317) (1996) - Vance B, Maier D.
- [Using Eels, A Practical Approach To Outerjoin And Antijoin Reordering](papers/join-order/using-eels,-a-practical-approach-to-outerjoin-and-antijoin-reordering.pdf) (2001) - Rao J, Lindsay B, Lohman G, et al.
- [Analysis Of Two Existing And One New Dynamic Programming Algorithm For The Generation Of Optimal Bushy Join Trees Without Cross Products](https://www.researchgate.net/profile/Thomas_Neumann2/publication/47861835_Analysis_of_Two_Existing_and_One_New_Dynamic_Programming_Algorithm_for_the_Generation_of_Optimal_Bushy_Join_Trees_without_Cross_Products/links/0912f506d90ad19031000000.pdff) (2006) - Moerkotte, Guido, and Thomas Neumann.
- [Optimal Top-Down Join Enumeration](https://dl.acm.org/doi/pdf/10.1145/1247480.1247567) (2007) - DeHaan D, Tompa F W.
- [Dynamic Programming Strikes Back](https://dl.acm.org/doi/pdf/10.1145/1376616.1376672) (2008) - Moerkotte, Guido, and Thomas Neumann.
- [On The Correct And Complete Enumeration Of The Core Search Space](https://dl.acm.org/doi/pdf/10.1145/2463676.2465314) (2013) - Moerkotte, Guido, Pit Fender, and Marius Eich.
- [How Good Are Query Optimizers, Really?](https://dl.acm.org/doi/pdf/10.14778/2850583.2850594) (2015) - Leis, Viktor, et al.
- [Improving Join Reorderability With Compensation Operators](https://dl.acm.org/doi/pdf/10.1145/3183713.3183731) (2018) - Wang, TaiNing, and Chee-Yong Chan.
- [Adaptive Optimization Of Very Large Join Queries](https://dl.acm.org/doi/pdf/10.1145/3183713.3183733) (2018) - Neumann, Thomas, and Bernhard Radke.

### Cost Model

- [Modelling Costs For A Mm-Dbms](https://www.semanticscholar.org/paper/Modelling-Costs-for-a-MM-DBMS-Listgarten-Neimat/42b88445cfb28fbe4b6539c97674a8fa9815e635) (1996) - Listgarten, Sherry, and Marie-Anne Neimat.
- [Seeking The Truth About Ad Hoc Join Costs](papers/cost-model/seeking-the-truth-about-ad-hoc-join-costs.pdf) (1997) - Haas, Laura M., et al.
- [Approximation Schemes For Many-Objective Query Optimization](https://dl.acm.org/doi/pdf/10.1145/2588555.2610527) (2014) - Trummer, Immanuel, and Christoph Koch.
- [Multi-Objective Parametric Query Optimization](https://dl.acm.org/doi/pdf/10.1145/3068612) (2015) - Trummer, Immanuel, and Christoph Koch.

### Statistics

- [Accurate Estimation Of The Number Of Tuples Satisfying A Condition](https://dl.acm.org/doi/pdf/10.1145/971697.602294) (1984) - Piatetsky-Shapiro, Gregory, and Charles Connell.
- [Optimal Histograms For Limiting Worst-Case Error Propagation In The Size Of Join Results](https://dl.acm.org/doi/pdf/10.1145/169725.169708) (1993) - Ioannidis, Yannis E., and Stavros Christodoulakis.
- [Universality Of Serial Histograms](papers/statistics/universality-of-serial-histograms.pdf) (1993) - Ioannidis, Yannis E.
- [Balancing Histogram Optimality And Practicality For Query Result Size Estimation](https://dl.acm.org/doi/pdf/10.1145/568271.223841) (1995) - Ioannidis, Yannis E., and Viswanath Poosala.
- [Improved Histograms For Selectivity Estimation Of Range Predicates](https://dl.acm.org/doi/pdf/10.1145/235968.233342) (1996) - Poosala, Viswanath, et al.
- [The History Of Histograms](papers/statistics/the-history-of-histograms.pdf) (2003) - Ioannidis, Yannis.
- [Automated Statistics Collection In Db2 Udb](papers/statistics/automated-statistics-collection-in-db2-udb.pdf) (2004) - Aboulnaga, Ashraf, et al.
- [Adaptive Query Processing In The Looking Glass](papers/statistics/adaptive-query-processing-in-the-looking-glass.pdf) (2005) - Babu, Shivnath, and Pedro Bizarro.
- [Optimizer Plan Change Management: Improved Stability And Performance In Oracle 11G](https://dl.acm.org/doi/pdf/10.14778/1454159.1454175) (2008) - Ziauddin, Mohamed, et al.
- [Histograms Reloaded: The Merits Of Bucket Diversity](https://dl.acm.org/doi/pdf/10.1145/1807167.1807239) (2010) - Kanne, Carl-Christian, and Guido Moerkotte.
- [Synopses For Massive Data: Samples, Histograms, Wavelets, Sketches](papers/statistics/synopses-for-massive-data:-samples,-histograms,-wavelets,-sketches.pdf) (2011) - Cormode, Graham, et al.
- [Exploiting Ordered Dictionaries To Efficiently Construct Histograms With Q-Error Guarantees In Sap Hana](https://dl.acm.org/doi/pdf/10.1145/2588555.2595629) (2014) - Moerkotte, Guido, et al.
- [Adaptive Statistics In Oracle 12C](https://dl.acm.org/doi/pdf/10.14778/3137765.3137785) (2017) - Chakkappen, Sunil, et al.

### Probabilistic Counting

- [Probabilistic counting algorithms for data base applications](https://algo.inria.fr/flajolet/Publications/FlMa85.pdf) (1985) - Flajolet, Philippe; Martin, G. Nigel.
- [Towards Estimation Error Guarantees For Distinct Values](https://dl.acm.org/doi/pdf/10.1145/335168.335230) (2000) - Charikar, Moses, et al.
- [Distinct Sampling For Highly-Accurate Answers To Distinct Values Queries And Event Reports](papers/probabilistic-counting/distinct-sampling-for-highly-accurate-answers-to-distinct-values-queries-and-event-reports.pdf) (2001) - Gibbons, Phillip B.
- [Leo – Db2’s Learning Optimizer](papers/probabilistic-counting/leo-–-db2’s-learning-optimizer.pdf) (2001) - Stillger, Michael, et al.
- [An Improved Data Stream Summary: The Count-Min Sketch And Its Applications, Journal Of Algorithms](papers/probabilistic-counting/an-improved-data-stream-summary:-the-count-min-sketch-and-its-applications,-journal-of-algorithms.pdf) (2005) - Cormode, Graham, and Shan Muthukrishnan.
- [New Estimation Algorithms For Streaming Data: Count-Min Can Do More](https://www.academia.edu/download/31052190/cmm.pdf) (2007) - Deng, Fan, and Davood Rafiei.
- [Preventing Bad Plans By Bounding The Impact Of Cardinality Estimation Errors](https://dl.acm.org/doi/pdf/10.14778/1687627.1687738) (2009) - Moerkotte, Guido, Thomas Neumann, and Gabriele Steidl.
- [Pessimistic Cardinality Estimation: Tighter Upper Bounds For Intermediate Join Cardinalities](https://dl.acm.org/doi/pdf/10.1145/3299869.3319894) (2019) - Cai, Walter, Magdalena Balazinska, and Dan Suciu.
- [Deep Unsupervised Cardinality Estimation](papers/probabilistic-counting/deep-unsupervised-cardinality-estimation.pdf) (2019) - Yang, Zongheng, et al.
- [Neurocard: One Cardinality Estimator For All Tables](papers/probabilistic-counting/neurocard:-one-cardinality-estimator-for-all-tables.pdf) (2020) - Yang, Zongheng, et al.

### Execution Engine

- [Querye Valuation Techniques For Large Databases](https://dl.acm.org/doi/pdf/10.1145/152610.152611) (1993) - Graefe G.
- [Volcano - An Extensible And Parallel Query Evaluation System](papers/execution-engine/volcano---an-extensible-and-parallel-query-evaluation-system.pdf) (1994) - Graefe, Goetz.
- [Monetdb/X100: Hyper-Pipelining Query Execution](https://www.researchgate.net/profile/Niels-Nes/publication/45338800_MonetDBX100_Hyper-Pipelining_Query_Execution/links/0deec520cd1e8a3607000000/MonetDB-X100-Hyper-Pipelining-Query-Execution.pdf) (2005) - Boncz, Peter A., Marcin Zukowski, and Niels Nes.
- [Efficiently Compiling Efficient Query Plans For Modern Hardware](https://dl.acm.org/doi/pdf/10.14778/2002938.2002940) (2011) - Neumann, Thomas.
- [Multi-Core, Main-Memory Joins: Sort Vs. Hash Revisited](https://dl.acm.org/doi/pdf/10.14778/2732219.2732227) (2013) - Balkesen, Cagri, et al.
- [Main-Memory Hash Joins On Modern Processor Architectures](papers/execution-engine/main-memory-hash-joins-on-modern-processor-architectures.pdf) (2014) - Balkesen Ç, Teubner J, Alonso G, et al.
- [Morsel-Driven Parallelism: A Numa-Aware Query Evaluation Framework For The Many-Core Age](https://dl.acm.org/doi/pdf/10.1145/2588555.2610507) (2014) - Leis, Viktor, et al.
- [Relaxed Operator Fusion For In-Memory Databases: Making Compilation, Vectorization, And Prefetching Work Together At Last](https://dl.acm.org/doi/pdf/10.14778/3151113.3151114) (2017) - Menon, Prashanth, Todd C. Mowry, and Andrew Pavlo.
- [Looking Ahead Makes Query Plans Robust](https://dl.acm.org/doi/pdf/10.14778/3090163.3090167) (2017) - Zhu, Jianqiao, et al.
- [Everything You Always Wanted To Know About Compiled And Vectorized Queries But Were Afraid To Ask](https://dl.acm.org/doi/pdf/10.14778/3275366.3284966) (2018) - Kersten, Timo, et al.
- [Surf: Practical Range Query Filtering With Fast Succinct Tries](https://dl.acm.org/doi/pdf/10.1145/3183713.3196931) (2018) - Zhang, Huanchen, et al.
- [Adaptive Execution Of Compiled Queries](https://15721.courses.cs.cmu.edu/spring2019/papers/19-compilation/kohn-icde2018.pdff) (2018) - Kohn, André, Viktor Leis, and Thomas Neumann.

### Parallel Execution

- [Db2 Parallel Edition](papers/parallel-execution/db2-parallel-edition.pdf) (1995) - Baru, Chaitanya K., et al.
- [Parallel Sql Execution In Oracle 10G](https://dl.acm.org/doi/pdf/10.1145/1007568.1007666) (2004) - Cruanes, Thierry, Benoit Dageville, and Bhaskar Ghosh.
- [Query Optimization In Microsoft Sql Server Pdw](https://dl.acm.org/doi/pdf/10.1145/2213836.2213953) (2012) - Shankar, Srinath, et al.
- [Adaptive And Big Data Scale Parallel Execution In Oracle](https://dl.acm.org/doi/pdf/10.14778/2536222.2536235) (2013) - Bellamkonda, Srikanth, et al.
- [Optimizing Queries Over Partitioned Tables In Mpp Systems](https://dl.acm.org/doi/pdf/10.1145/2588555.2595640) (2014) - Antova, Lyublena, et al.

## Storage Engine

### Storage Media

- [The 5 Minute Rule For Trading Memory For Disc Accesses And The 5 Byte Rule For Trading Memory For Cpu Time](https://dl.acm.org/doi/pdf/10.1145/38713.38755) (1987) - Gray, Jim, and Franco Putzolu.
- [The Five-Minute Rule Ten Years Later, And Other Computer Storage Rules Of Thumb](https://dl.acm.org/doi/pdf/10.1145/271074.271094) (1997) - Gray, Jim, and Goetz Graefe.
- [The Five Minute Rule 20 Years Later And How Flash Memory Changes The Rules](https://dl.acm.org/doi/pdf/10.1145/1363189.1363198) (2008) - Graefe, Goetz.
- [The Five Minute Rule Thirty Years Later And Its Impact On The Storage Hierarchy](papers/storage-media/the-five-minute-rule-thirty-years-later-and-its-impact-on-the-storage-hierarchy.pdf) (2017) - Appuswamy, Raja, et al.

### Storage Structure

- [The Ubiquitous B-Tree](https://dl.acm.org/doi/pdf/10.1145/356770.356776) (1979) - Comer, Douglas.
- [Principles Of Database Buffer Management](https://dl.acm.org/doi/pdf/10.1145/1994.2022) (1984) - Effelsberg W, Haerder T.
- [The Log-Structured Merge-Tree (Lsm-Tree)](papers/storage-structure/the-log-structured-merge-tree-(lsm-tree).pdf) (1996) - O’Neil, Patrick, et al.
- [A Comparison Of Fractal Trees To Log-Structured Merge (Lsm) Trees](papers/storage-structure/a-comparison-of-fractal-trees-to-log-structured-merge-(lsm)-trees.pdf) (2014) - Kuszmaul, Bradley C.
- [Design Tradeoffs Of Data Access Methods](https://dl.acm.org/doi/pdf/10.1145/2882903.2912569) (2016) - Athanassoulis, Manos, and Stratos Idreos.
- [Designing Access Methods: The Rum Conjecture](papers/storage-structure/designing-access-methods:-the-rum-conjecture.pdf) (2016) - Athanassoulis, Manos, et al.
- [Wisckey: Separating Keys From Values In Ssd-Conscious Storage](https://dl.acm.org/doi/pdf/10.1145/3033273) (2017) - Lu, Lanyue, et al.
- [Managing Non-Volatile Memory In Database Systems](https://dl.acm.org/doi/pdf/10.1145/3183713.3196897) (2018) - van Renen, Alexander, et al.
- [Leanstore: In-Memory Data Management Beyond Main Memory](papers/storage-structure/leanstore:-in-memory-data-management-beyond-main-memory.pdf) (2018) - Leis, Viktor, et al.
- [The Case For Learned Index Structures](https://dl.acm.org/doi/pdf/10.1145/3183713.3196909) (2018) - Kraska, Tim, et al.
- [Lsm-Based Storage Techniques: A Survey](papers/storage-structure/lsm-based-storage-techniques:-a-survey.pdf) (2019) - Luo, Chen, and Michael J. Carey.
- [Learning Multi-Dimensional Indexes](https://dl.acm.org/doi/pdf/10.1145/3318464.3380579) (2019) - Nathan, Vikram, et al.
- [Umbra: A Disk-Based System With In-Memory Performance](papers/storage-structure/umbra:-a-disk-based-system-with-in-memory-performance.pdf) (2020) - Neumann, Thomas, and Michael J. Freitag.
- [Xindex: A Scalable Learned Index For Multicore Data Storage](https://dl.acm.org/doi/pdf/10.1145/3332466.3374547) (2020) - Tang, Chuzhe, et al.
- [The Pgm-Index: A Fully-Dynamic Compressed Learned Index With Provable Worst-Case Bounds](https://dl.acm.org/doi/pdf/10.14778/3389133.3389135) (2020) - Ferragina, Paolo, and Giorgio Vinciguerra.
- [From Wisckey To Bourbon: A Learned Index For Log-Structured Merge Trees](papers/storage-structure/from-wisckey-to-bourbon:-a-learned-index-for-log-structured-merge-trees.pdf) (2020) - Dai, Yifan, et al.
- [Caas-Lsm: Compaction-As-A-Service For Lsm-Based Key-Value Stores In Storage Disaggregated Infrastructure](papers/storage-structure/caas-lsm:-compaction-as-a-service-for-lsm-based-key-value-stores-in-storage-disaggregated-infrastructure.pdf) (2024) - Yu, Qiaolin et al.

### Transaction

- [The Notions Of Consistency And Predicate Locks In A Database System](https://dl.acm.org/doi/pdf/10.1145/360363.360369) (1976) - Eswaran, Kapali P., et al.
- [Concurrency Control In Distributed Database Systems](https://dl.acm.org/doi/pdf/10.1145/356842.356846) (1981) - Bernstein, Philip A., and Nathan Goodman.
- [On Optimistic Methods For Concurrency Control](https://dl.acm.org/doi/pdf/10.1145/319566.319567) (1981) - Kung, Hsiang-Tsung, and John T. Robinson.
- [Principles Of Transaction-Oriented Database Recovery](https://dl.acm.org/doi/10.1145/289.291) (1983) - Haerder, Theo, and Andreas Reuter.
- [Multiversion Concurrency Control - Theory And Algorithms](https://dl.acm.org/doi/pdf/10.1145/319996.319998) (1983) - Bernstein, Philip A., and Nathan Goodman.
- [Aries: A Transaction Recovery Method Supporting Fine-Granularity Locking And Partial Rollbacks Using Write-Ahead Logging](https://dl.acm.org/doi/pdf/10.1145/128765.128770) (1992) - Mohan C, Haderle D, Lindsay B, et al.
- [A Critique Of Ansi Sql Isolation Levels](https://dl.acm.org/doi/pdf/10.1145/568271.223785) (1995) - Berenson, Hal, et al.
- [Generalized Isolation Level Definitions](papers/transaction/generalized-isolation-level-definitions.pdf) (2000) - Adya, Atul, Barbara Liskov, and Patrick O'Neil.
- [Large-Scale Incremental Processing Using Distributed Transactions And Notifications](papers/transaction/large-scale-incremental-processing-using-distributed-transactions-and-notifications.pdf) (2010) - Peng D, Dabek F.
- [Serializable Snapshot Isolation In Postgresql](https://arxiv.org/pdf/1208.4179.pdf,) (2012) - Ports, Dan RK, and Kevin Grittner.
- [Calvin: Fast Distributed Transactions For Partitioned Database Systems](https://dl.acm.org/doi/pdf/10.1145/2213836.2213838) (2012) - Thomson, Alexander, et al.
- [Maat: Effective And Scalable Coordination Of Distributed Transactions In The Cloud](https://dl.acm.org/doi/pdf/10.14778/2732269.2732270) (2014) - Mahmoud, Hatem A., et al.
- [Staring Into The Abyss: An Evaluation Of Concurrency Control With One Thousand Cores](papers/transaction/staring-into-the-abyss:-an-evaluation-of-concurrency-control-with-one-thousand-cores.pdf) (2014) - Yu, Xiangyao, et al.
- [An Evaluation Of The Advantages And Disadvantages Of Deterministic Database Systems](https://dl.acm.org/doi/pdf/10.14778/2732951.2732955) (2014) - Ren, Kun, Alexander Thomson, and Daniel J. Abadi.
- [Fast Serializable Multi-Version Concurrency Control For Main-Memory Database Systems](https://dl.acm.org/doi/pdf/10.1145/2723372.2749436) (2015) - Neumann, Thomas, Tobias Mühlbauer, and Alfons Kemper.
- [An Empirical Evaluation Of In-Memory Multi-Version Concurrency Control](https://dl.acm.org/doi/pdf/10.14778/3067421.3067427) (2017) - Wu, Yingjun, et al.
- [An Evaluation Of Distributed Concurrency Control](https://dl.acm.org/doi/pdf/10.14778/3055540.3055548) (2017) - Harding, Rachael, et al.
- [Scalable Garbage Collection For In-Memory Mvcc Systems](https://dl.acm.org/doi/pdf/10.14778/3364324.3364328) (2019) - Böttcher, Jan, et al.

### Scheduling

- [Automated Demand-Driven Resource Scaling In Relational Database-As-A-Service](https://dl.acm.org/doi/pdf/10.1145/2882903.2903733) (2016) - Das, Sudipto, et al.
- [Autoscaling Tiered Cloud Storage In Anna](https://dl.acm.org/doi/pdf/10.14778/3311880.3311881) (2019) - Wu, Chenggang, Vikram Sreekanti, and Joseph M. Hellerstein.
- [Adaptive Htap Through Elastic Resource Scheduling](https://dl.acm.org/doi/pdf/10.1145/3318464.3389783) (2020) - Raza, Aunn, et al.
- [Morphosys: Automatic Physical Design Metamorphosis For Distributed Database Systems](https://dl.acm.org/doi/pdf/10.14778/3424573.3424578) (2020) - Abebe, Michael, Brad Glasbergen, and Khuzaima Daudjee.

## Miscellaneous

### Workload

- [Tpc-H Analyzed: Hidden Messages And Lessons Learned From An Influential Benchmark](https://www.researchgate.net/profile/Peter-Boncz/publication/291257517_TPC-H_Analyzed_Hidden_Messages_and_Lessons_Learned_from_an_Influential_Benchmark/links/5852dbf708ae95fd8e1d749b/TPC-H-Analyzed-Hidden-Messages-and-Lessons-Learned-from-an-Influential-Benchmark.pdff) (2013) - Boncz, Peter, Thomas Neumann, and Orri Erling.
- [Quantifying Tpch Choke Points And Their Optimizations](https://dl.acm.org/doi/pdf/10.14778/3389133.3389138) (2020) - Dreseler, Markus, et al.

### Network

- [The End Of Slow Networks: It's Time For A Redesign](papers/network/the-end-of-slow-networks:-it's-time-for-a-redesign.pdf) (2015) - Binnig, Carsten, et al.
- [Accelerating Relational Databases By Leveraging Remote Memory And Rdma](https://dl.acm.org/doi/pdf/10.1145/2882903.2882949) (2016) - Li, Feng, et al.
- [Don't Hold My Data Hostage: A Case For Client Protocol Redesign](https://dl.acm.org/doi/pdf/10.14778/3115404.3115408) (2017) - Raasveldt, Mark, and Hannes Mühleisen.

### Quality

- [Testing The Accuracy Of Query Optimizers](https://dl.acm.org/doi/pdf/10.1145/2304510.2304525) (2012) - Gu, Zhongxian, Mohamed A. Soliman, and Florian M.

### Diagnosis and Tuning

- [Automatic Sql Tuning In Oracle 10G](papers/diagnosis-and-tuning/automatic-sql-tuning-in-oracle-10g.pdf) (2004) - Dageville B, Das D, Dias K, et al.
- [Automatic Performance Diagnosis And Tuning In Oracle](papers/diagnosis-and-tuning/automatic-performance-diagnosis-and-tuning-in-oracle.pdf) (2005) - Dias K, Ramacher M, Shaft U, et al.

