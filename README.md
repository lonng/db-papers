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
- [Miscellaneous](#miscellaneous)
  - [Workload](#workload)
  - [Network](#network)
  - [Quality](#quality)
  - [Diagnosis and Tuning](#diagnosis-and-tuning)

## Basics

### Essentials

- [A relational model of data for large shared data banks](https://dl.acm.org/doi/pdf/10.1145/362384.362685) (1970) - Codd, Edgar F.
- [SEQUEL: A structured English query language](https://dl.acm.org/doi/pdf/10.1145/800296.811515) (1974) - Chamberlin, Donald D., and Raymond F. Boyce.
- [INGRES: a relational data base system](https://dl.acm.org/doi/pdf/10.1145/1499949.1500029) (1975) - Held, G. D., M. R. Stonebraker, and Eugene Wong.
- [Extending the database relational model to capture more meaning](https://dl.acm.org/doi/pdf/10.1145/320107.320109) (1979) - Codd, Edgar F.
- [A critique of the SQL database language](https://dl.acm.org/doi/pdf/10.1145/984549.984551) (1984) - Date, C. J.

### Consensus

- [The Part-Time Parliament](https://dl.acm.org/doi/pdf/10.1145/3335772.3335939) (1998) - Lamport, Leslie.
- [Paxos Made Simple](https://www.microsoft.com/en-us/research/publication/2016/12/paxos-simple-Copy.pdf) (2001) - Lamport, Leslie.
- [Consensus: Bridging theory and practice](https://web.stanford.edu/~ouster/cgi-bin/papers/OngaroPhD.pdf) (2014) - Ongaro, Diego.
- [In search of an understandable consensus algorithm (extended version)](http://pages.cs.wisc.edu/~remzi/Classes/739/Spring2004/Papers/raft.pdf) (2014) - Ongaro, Diego, and John Ousterhout.
- [Distributed consensus revised](https://www.repository.cam.ac.uk/bitstream/handle/1810/291682/thesis.pdf?sequence=1) (2019) - Howard, Heidi.
- [A Generalised Solution to Distributed Consensus](https://arxiv.org/pdf/1902.06776) (2019) - Howard, Heidi, and Richard Mortier.
- [Paxos vs Raft: Have we reached consensus on distributed consensus?](https://dl.acm.org/doi/pdf/10.1145/3380787.3393681) (2020) - Howard, Heidi, and Richard Mortier.

### Consistency

- [Consistency Tradeoffs in Modern Distributed Database System Design](https://www.cs.umd.edu/~abadi/papers/abadi-pacelc.pdf) (2012) - Abadi, Daniel.
- [Logical physical clocks and consistent snapshots in globally distributed databases](https://cse.buffalo.edu/tech-reports/2014-04.pdf) (2014) - Kulkarni S S, Demirbas M, Madappa D, et al.
- [Ark: A Real-World Consensus Implementation](https://arxiv.org/pdf/1407.4765) (2014) - Kasheff, Zardosht, and Leif Walsh.
- [PolarFS: an ultra-low latency and failure resilient distributed file system for shared storage cloud database](https://dl.acm.org/doi/pdf/10.14778/3229863.3229872) (2018) - Cao, Wei, et al.
- [Anna: A kvs for any scale](https://www2.eecs.berkeley.edu/Pubs/TechRpts/2019/EECS-2019-122.pdf) (2018) - Wu, Chenggang, et al.
- [Strong and efficient consistency with consistency-aware durability](https://dl.acm.org/doi/pdf/10.1145/3423138) (2021) - Ganesan, Aishwarya, et al.

## System Design

### RDBMS

- [System R: Relational Approach to Database Management](https://dl.acm.org/doi/pdf/10.1145/320455.320457) (1976) - Astrahan, Morton M., et al.
- [The design and implementation of INGRES](https://dl.acm.org/doi/10.1145/320473.320476) (1976) - Stonebraker, Michael, et al.
- [The design of Postgres](https://dl.acm.org/doi/pdf/10.1145/16856.16888) (1986) - Stonebraker, Michael, and Lawrence A. Rowe.
- [Query Processing in Main Memory Database Management Systems](https://dl.acm.org/doi/pdf/10.1145/16894.16878) (1986) - Lehman, Tobin J., and Michael J. Carey.
- [Megastore: Providing Scalable, Highly Available Storage for Interactive Services](http://pages.cs.wisc.edu/~akella/CS838/F12/838-CloudPapers/Megastore.pdf) (2011) - Baker J, Bond C, Corbett J C, et al.
- [Spanner: Google's globally distributed database](https://dl.acm.org/doi/pdf/10.1145/2491245) (2013) - Corbett, James C., et al.
- [Online, Asynchronous Schema Change in F1](https://dl.acm.org/doi/pdf/10.14778/2536222.2536230) (2013) - Rae, Ian, et al.
- [Amazon aurora: Design considerations for high throughput cloud-native relational databases](https://dl.acm.org/doi/pdf/10.1145/3035918.3056101) (2017) - Verbitski, Alexandre, et al.
- [Looking Back at Postgres](https://arxiv.org/pdf/1901.01973) (2019) - Hellerstein, Joseph M.
- [CockroachDB: The Resilient Geo-Distributed SQL Database](https://dl.acm.org/doi/pdf/10.1145/3318464.3386134) (2020) - Taft, Rebecca, et al.
- [F1 Lightning: HTAP as a Service](https://dl.acm.org/doi/pdf/10.14778/3415478.3415553) (2020) - Yang, Jiacheng, et al.
- [TiDB: a Raft-based HTAP database](https://dl.acm.org/doi/pdf/10.14778/3415478.3415535) (2020) - Huang, Dongxu, et al.
- [PolarDB Serverless: A Cloud Native Database for Disaggregated Data Centers](https://dl.acm.org/doi/pdf/10.1145/3448016.3457560) (2021) - Cao, Wei, et al.

### NoSQL

- [Bigtable: A Distributed Storage System for Structured Data](https://dl.acm.org/doi/pdf/10.1145/1365815.1365816) (2006) - Chang, Fay, et al.
- [Dynamo: Amazon’s Highly Available Key-value Store](https://dl.acm.org/doi/pdf/10.1145/1323293.1294281) (2007) - DeCandia, Giuseppe, et al.
- [PNUTS: Yahoo!’s Hosted Data Serving Platform](https://dl.acm.org/doi/pdf/10.14778/1454159.1454167) (2008) - Cooper, Brian F., et al.
- [Cassandra - A Decentralized Structured Storage System](https://dl.acm.org/doi/pdf/10.1145/1773912.1773922) (2010) - Lakshman, Avinash, and Prashant Malik.
- [Windows azure storage: a highly available cloud storage service with strong consistency](https://dl.acm.org/doi/pdf/10.1145/2043556.2043571) (2011) - Calder, Brad, et al.
- [Azure data lake store: a hyperscale distributed file service for big data analytics](https://dl.acm.org/doi/pdf/10.1145/3035918.3056100) (2017) - Ramakrishnan, Raghu, et al.
- [PNUTS to Sherpa: Lessons from Yahoo!’s Cloud Database](https://dl.acm.org/doi/pdf/10.14778/3352063.3352146) (2019) - Cooper, Brian F., et al.

## SQL Engine

### Optimizer Framework

- [Access Path Selection in a Relational Database Management System](https://dl.acm.org/doi/pdf/10.1145/582095.582099) (1979) - Selinger, P. Griffiths, et al.
- [Query Optimization by Simulated Annealing](https://dl.acm.org/doi/pdf/10.1145/38713.38722) (1987) - Ioannidis, Yannis E., and Eugene Wong.
- [The EXODUS Optimizer Generator](https://dl.acm.org/doi/pdf/10.1145/38713.38734) (1987) - Graefe, Goetz, and David J. DeWitt.
- [Extensible/Rule Based Query Rewrite Optimization in Starburst](https://dl.acm.org/doi/pdf/10.1145/141484.130294) (1992) - Pirahesh, Hamid, Joseph M. Hellerstein, and Waqar Hasan.
- [The Volcano Optimizer Generator- Extensibility and Efficient Search](https://www.cse.iitb.ac.in/infolab/Data/Courses/CS632/Papers/Volcano-graefe.pdf) (1993) - Graefe, Goetz, and William J. McKenna.
- [The Cascades Framework for Query Optimization](https://liuyehcf.github.io/resources/paper/The-Cascades-Framework-For-Query-Optimization.pdf) (1995) - Graefe, Goetz.
- [An Overview of Query Optimization in Relational Systems](https://dl.acm.org/doi/pdf/10.1145/275487.275492) (1998) - Chaudhuri, Surajit.
- [Robust Query Processing through Progressive Optimization](https://dl.acm.org/doi/pdf/10.1145/1007568.1007642) (2004) - Markl, Volker, et al.
- [Orca: A Modular Query Optimizer Architecture for Big Data](https://dl.acm.org/doi/pdf/10.1145/2588555.2595637) (2014) - Soliman, Mohamed A., et al.
- [Parallelizing Query Optimization on Shared-Nothing Architectures](https://arxiv.org/pdf/1511.01768) (2015) - Trummer, Immanuel, and Christoph Koch.
- [The MemSQL Query Optimizer: A modern optimizer for real-time analytics in a distributed database](https://dl.acm.org/doi/pdf/10.14778/3007263.3007277) (2016) - Chen, Jack, et al.

### Transformation

- [Processing queries with quantifiers a horticultural approach](https://dl.acm.org/doi/pdf/10.1145/588058.588075) (1983) - Dayal, Umeshwar.
- [Translating SQL into relational algebra: Optimization, semantics, and equivalence of SQL queries](https://www.academia.edu/download/50687636/tse.1985.23222320161202-29901-8u86ef.pdf) (1985) - Ceri, Stefano, and Georg Gottlob.
- [Grammar-like Functional Rules for Representing Query Optimization Alternatives,](https://dl.acm.org/doi/pdf/10.1145/971701.50204) (1988) - Lohman, Guy M.
- [Query Optimization by Predicate Move-Around](https://www.researchgate.net/profile/Inderpal-Mumick/publication/2754592_Query_Optimization_by_Predicate_Move-Around/links/0f317534d437e49755000000/Query-Optimization-by-Predicate-Move-Around.pdf) (1994) - Levy, Alon Y., Inderpal Singh Mumick, and Yehoshua Sagiv.
- [Eager Aggregation and Lazy Aggregation](https://www.researchgate.net/profile/Per-Ake-Larson/publication/2733082_Eager_Aggregation_and_Lazy_Aggregation/links/02bfe50ce6de3dad7c000000/Eager-Aggregation-and-Lazy-Aggregation.pdf) (1995) - Yan, Weipeng P., and Per-Bike Larson.
- [Parameterized Queries and Nesting Equivalences](https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/tr-2000-31.pdf) (2000) - Galindo-Legaria, C. A.
- [Cost-based query transformation in Oracle](https://www.researchgate.net/profile/Rafi-Ahmed-2/publication/221311318_Cost-Based_Query_Transformation_in_Oracle/links/572bbc5e08aef7c7e2c6b829/Cost-Based-Query-Transformation-in-Oracle.pdf) (2006) - Ahmed, Rafi, et al.

### Nested Query

- [Using semi-joins to solve relational queries](https://dl.acm.org/doi/pdf/10.1145/322234.322238) (1981) - Bernstein, Philip A., and Dah-Ming W. Chiu.
- [On optimizing an SQL-like nested query](https://dl.acm.org/doi/pdf/10.1145/319732.319745) (1982) - Kim, Won.
- [Optimization of nested queries in a distributed relational database](https://citeseerx.ist.psu.edu/document?repid=rep1&type=pdf&doi=12fd1fe22687f5944613832de4e64ef902043aec) (1984) - L&man, Guy M., et al.
- [SQL-like and Quel-like correlation queries with aggregates revisited](http://www2.eecs.berkeley.edu/Pubs/TechRpts/1984/ERL-m-84-75.pdf) (1984) - Kiessling, Werner.
- [Translating SQL into relational algebra: Optimization, semantics, and equivalence of SQL queries](https://www.academia.edu/download/50687636/tse.1985.23222320161202-29901-8u86ef.pdf) (1985) - Ceri, Stefano, and Georg Gottlob.
- [Optimization of nested SQL queries revisited](https://dl.acm.org/doi/pdf/10.1145/38714.38723) (1987) - Ganski, Richard A., and Harry KT Wong.
- [A Unitied Approach to Processing Queries That Contain Nested Subqueries, Aggregates, and Quantifiers](https://vldb.org/conf/1987/P197.PDF) (1987) - Dayal, Umeshwar.
- [Optimization of correlated SQL queries in a relational database management system](https://patentimages.storage.googleapis.com/3b/24/39/a947424a6eb0ea/US5822750.pdf) (1998) - Jou, Michelle M., Ting Yu Leung, and Mir Hamid Pirahesh.
- [Orthogonal Optimization of Subqueries and Aggregation](https://dl.acm.org/doi/pdf/10.1145/376284.375748) (2001) - Galindo-Legaria, César, and Milind Joshi.
- [WinMagic : Subquery Elimination Using Window Aggregation](https://dl.acm.org/doi/pdf/10.1145/872757.872840) (2003) - Zuzarte, Calisto, et al.
- [Execution strategies for SQL subqueries](https://dl.acm.org/doi/pdf/10.1145/1247480.1247598) (2007) - Elhemali, Mostafa, et al.
- [Enhanced subquery optimizations in Oracle](https://dl.acm.org/doi/pdf/10.14778/1687553.1687563) (2009) - Bellamkonda, Srikanth, et al.
- [Unnesting Arbitrary Queries](https://dl.gi.de/bitstream/handle/20.500.12116/2418/383.pdf?sequence=1)) (2015) - Neumann, Thomas, and Alfons Kemper.

### Functional Dependencies

- [Fundamental Techniques for Order Optimization](https://dl.acm.org/doi/pdf/10.1145/233269.233320) (1996) - Simmen, David, Eugene Shekita, and Timothy Malkemus.
- [[Thesis] Exploiting Functional Dependence in Query Optimization](https://cs.uwaterloo.ca/research/tr/2000/11/CS-2000-11.thesis.pdf) (2000) - Paulley, Glenn Norman.
- [An Efficient Framework for Order Optimization](https://madoc.bib.uni-mannheim.de/736/1/TR-03-011.pdf) (2004) - Neumann, Thomas, and Guido Moerkotte.
- [Incorporating Partitioning and Parallel Plans into the SCOPE Optimizer](http://www.cs.albany.edu/~jhh/courses/readings/zhou10.pdf) (2010) - Zhou, Jingren, Per-Ake Larson, and Ronnie Chaiken.
- [Accelerating Queries with GroupBy and Join by Group join](https://dl.acm.org/doi/pdf/10.14778/3402707.3402723) (2011) - Moerkotte, Guido, and Thomas Neumann.

### Join Order

- [Access paths in the" Abe" statistical query facility](https://dl.acm.org/doi/pdf/10.1145/582353.582382) (1982) - Klug, Anthony.
- [Extending the Algebraic Framework of Query Processing to Handle Outerjoins](https://www.vldb.org/conf/1984/P334.PDF) (1984) - RosenthaI, A., and D. Reiner.
- [Analysis of Two Existing and One New Dynamic Programming Algorithm for the Generation of Optimal Bushy Join Trees without Cross Products](https://www.researchgate.net/profile/Thomas_Neumann2/publication/47861835_Analysis_of_Two_Existing_and_One_New_Dynamic_Programming_Algorithm_for_the_Generation_of_Optimal_Bushy_Join_Trees_without_Cross_Products/links/0912f506d90ad19031000000.pdff) (2006) - Moerkotte, Guido, and Thomas Neumann.
- [Dynamic programming strikes back](https://dl.acm.org/doi/pdf/10.1145/1376616.1376672) (2008) - Moerkotte, Guido, and Thomas Neumann.
- [On the Correct and Complete Enumeration of the Core Search Space](https://dl.acm.org/doi/pdf/10.1145/2463676.2465314) (2013) - Moerkotte, Guido, Pit Fender, and Marius Eich.
- [How Good Are Query Optimizers, Really?](https://dl.acm.org/doi/pdf/10.14778/2850583.2850594) (2015) - Leis, Viktor, et al.
- [The Complete Story of Joins](https://dl.gi.de/bitstreams/535a5d94-043d-4b1a-9062-fbaf8ed35468/download) (2017) - Neumann, Thomas, Viktor Leis, and Alfons Kemper.
- [Improving Join Reorderability with Compensation Operators](https://dl.acm.org/doi/pdf/10.1145/3183713.3183731) (2018) - Wang, TaiNing, and Chee-Yong Chan.
- [Adaptive Optimization of Very Large Join Queries](https://dl.acm.org/doi/pdf/10.1145/3183713.3183733) (2018) - Neumann, Thomas, and Bernhard Radke.

### Cost Model

- [Modelling Costs for a MM-DBMS](https://www.semanticscholar.org/paper/Modelling-Costs-for-a-MM-DBMS-Listgarten-Neimat/42b88445cfb28fbe4b6539c97674a8fa9815e635) (1996) - Listgarten, Sherry, and Marie-Anne Neimat.
- [SEEKing the truth about ad hoc join costs](https://minds.wisconsin.edu/bitstream/handle/1793/59726/TR1148.pdf?sequence=11) (1997) - Haas, Laura M., et al.
- [Approximation Schemes for Many-Objective Query Optimization](https://dl.acm.org/doi/pdf/10.1145/2588555.2610527) (2014) - Trummer, Immanuel, and Christoph Koch.
- [Multi-Objective Parametric Query Optimization](https://dl.acm.org/doi/pdf/10.1145/3068612) (2015) - Trummer, Immanuel, and Christoph Koch.

### Statistics

- [Accurate Estimation of the Number of Tuples Satisfying a Condition](https://dl.acm.org/doi/pdf/10.1145/971697.602294) (1984) - Piatetsky-Shapiro, Gregory, and Charles Connell.
- [Optimal Histograms for Limiting Worst-Case Error Propagation in the Size of Join Results](https://dl.acm.org/doi/pdf/10.1145/169725.169708) (1993) - Ioannidis, Yannis E., and Stavros Christodoulakis.
- [Universality of Serial Histograms](https://vldb.org/conf/1993/P256.PDF) (1993) - Ioannidis, Yannis E.
- [Balancing Histogram Optimality and Practicality for Query Result Size Estimation](https://dl.acm.org/doi/pdf/10.1145/568271.223841) (1995) - Ioannidis, Yannis E., and Viswanath Poosala.
- [Improved Histograms for Selectivity Estimation of Range Predicates](https://dl.acm.org/doi/pdf/10.1145/235968.233342) (1996) - Poosala, Viswanath, et al.
- [The History of Histograms](http://www.vldb.org/conf/2003/papers/S02P01.pdf) (2003) - Ioannidis, Yannis.
- [Automated Statistics Collection in DB2 UDB](http://www.vldb.org/conf/2004/IND5P3.PDF) (2004) - Aboulnaga, Ashraf, et al.
- [Adaptive Query Processing in the Looking Glass](https://eden.dei.uc.pt/~bizarro/papers/cidr2005_aqp.pdf) (2005) - Babu, Shivnath, and Pedro Bizarro.
- [Optimizer plan change management: improved stability and performance in Oracle 11g](https://dl.acm.org/doi/pdf/10.14778/1454159.1454175) (2008) - Ziauddin, Mohamed, et al.
- [Histograms Reloaded: The Merits of Bucket Diversity](https://dl.acm.org/doi/pdf/10.1145/1807167.1807239) (2010) - Kanne, Carl-Christian, and Guido Moerkotte.
- [Synopses for Massive Data: Samples, Histograms, Wavelets, Sketches](https://www.nowpublishers.com/article/DownloadSummary/DBS-004) (2011) - Cormode, Graham, et al.
- [Exploiting Ordered Dictionaries to Efficiently Construct Histograms with Q-Error Guarantees in SAP HANA](https://dl.acm.org/doi/pdf/10.1145/2588555.2595629) (2014) - Moerkotte, Guido, et al.
- [Adaptive Statistics in Oracle 12c](https://dl.acm.org/doi/pdf/10.14778/3137765.3137785) (2017) - Chakkappen, Sunil, et al.

### Probabilistic Counting

- [Towards Estimation Error Guarantees for Distinct Values](https://dl.acm.org/doi/pdf/10.1145/335168.335230) (2000) - Charikar, Moses, et al.
- [Distinct Sampling for Highly-Accurate Answers to Distinct Values Queries and Event Reports](http://www.vldb.org/conf/2001/P541.pdf) (2001) - Gibbons, Phillip B.
- [LEO – DB2’s LEarning Optimizer](http://www.vldb.org/conf/2001/P019.pdf) (2001) - Stillger, Michael, et al.
- [An Improved Data Stream Summary: The Count-Min Sketch and its Applications, Journal of Algorithms](http://twiki.di.uniroma1.it/pub/Ing_algo/WebHome/p14_Cormode_JAl_05.pdf) (2005) - Cormode, Graham, and Shan Muthukrishnan.
- [New Estimation Algorithms for Streaming Data: Count-min Can Do More](https://www.academia.edu/download/31052190/cmm.pdf) (2007) - Deng, Fan, and Davood Rafiei.
- [Preventing Bad Plans by Bounding the Impact of Cardinality Estimation Errors](https://dl.acm.org/doi/pdf/10.14778/1687627.1687738) (2009) - Moerkotte, Guido, Thomas Neumann, and Gabriele Steidl.
- [Pessimistic Cardinality Estimation: Tighter Upper Bounds for Intermediate Join Cardinalities](https://dl.acm.org/doi/pdf/10.1145/3299869.3319894) (2019) - Cai, Walter, Magdalena Balazinska, and Dan Suciu.
- [Deep Unsupervised Cardinality Estimation](https://arxiv.org/pdf/1905.04278) (2019) - Yang, Zongheng, et al.
- [NeuroCard: One Cardinality Estimator for All Tables](https://arxiv.org/pdf/2006.08109) (2020) - Yang, Zongheng, et al.

### Execution Engine

- [QueryEvaluationTechniquesfor LargeDatabas](https://dl.acm.org/doi/pdf/10.1145/152610.152611) (1993) - Graefe G.
- [Volcano - An Extensible and Parallel Query Evaluation System](https://15721.courses.cs.cmu.edu/spring2016/papers/graefe-ieee1994.pdf) (1994) - Graefe, Goetz.
- [MonetDB/X100: Hyper-Pipelining Query Execution](https://www.researchgate.net/profile/Niels-Nes/publication/45338800_MonetDBX100_Hyper-Pipelining_Query_Execution/links/0deec520cd1e8a3607000000/MonetDB-X100-Hyper-Pipelining-Query-Execution.pdf) (2005) - Boncz, Peter A., Marcin Zukowski, and Niels Nes.
- [Efficiently Compiling Efficient Query Plans for Modern Hardware](https://dl.acm.org/doi/pdf/10.14778/2002938.2002940) (2011) - Neumann, Thomas.
- [Multi-Core, Main-Memory Joins: Sort vs. Hash Revisited](https://dl.acm.org/doi/pdf/10.14778/2732219.2732227) (2013) - Balkesen, Cagri, et al.
- [Morsel-Driven Parallelism: A NUMA-Aware Query Evaluation Framework for the Many-Core Age](https://dl.acm.org/doi/pdf/10.1145/2588555.2610507) (2014) - Leis, Viktor, et al.
- [Relaxed Operator Fusion for In-Memory Databases: Making Compilation, Vectorization, and Prefetching Work Together At Last](https://dl.acm.org/doi/pdf/10.14778/3151113.3151114) (2017) - Menon, Prashanth, Todd C. Mowry, and Andrew Pavlo.
- [Looking Ahead Makes Query Plans Robust](https://dl.acm.org/doi/pdf/10.14778/3090163.3090167) (2017) - Zhu, Jianqiao, et al.
- [Everything You Always Wanted to Know About Compiled and Vectorized Queries But Were Afraid to Ask](https://dl.acm.org/doi/pdf/10.14778/3275366.3284966) (2018) - Kersten, Timo, et al.
- [SuRF: Practical Range Query Filtering with Fast Succinct Tries](https://dl.acm.org/doi/pdf/10.1145/3183713.3196931) (2018) - Zhang, Huanchen, et al.
- [Adaptive Execution of Compiled Queries](https://15721.courses.cs.cmu.edu/spring2019/papers/19-compilation/kohn-icde2018.pdff) (2018) - Kohn, André, Viktor Leis, and Thomas Neumann.

### MPP Optimizations

- [DB2 Parallel Edition](https://grape.ics.uci.edu/wiki/asterix/raw-attachment/wiki/cs295-2009-fall/ParallelDB2.pdf) (1995) - Baru, Chaitanya K., et al.
- [Parallel SQL execution in Oracle 10g](https://dl.acm.org/doi/pdf/10.1145/1007568.1007666) (2004) - Cruanes, Thierry, Benoit Dageville, and Bhaskar Ghosh.
- [Query Optimization in Microsoft SQL Server PDW](https://dl.acm.org/doi/pdf/10.1145/2213836.2213953) (2012) - Shankar, Srinath, et al.
- [Adaptive and big data scale parallel execution in Oracle](https://dl.acm.org/doi/pdf/10.14778/2536222.2536235) (2013) - Bellamkonda, Srikanth, et al.
- [Optimizing Queries over Partitioned Tables in MPP Systems](https://dl.acm.org/doi/pdf/10.1145/2588555.2595640) (2014) - Antova, Lyublena, et al.

## Storage Engine

### Storage Structure

- [The Ubiquitous B-Tree](https://dl.acm.org/doi/pdf/10.1145/356770.356776) (1979) - Comer, Douglas.
- [The 5 Minute Rule for Trading Memory for Disc Accesses and the 5 Byte Rule for Trading Memory for CPU Time](https://dl.acm.org/doi/pdf/10.1145/38713.38755) (1987) - Gray, Jim, and Franco Putzolu.
- [The Log-Structured Merge-Tree (LSM-Tree)](https://www.inf.ufpr.br/eduardo/ensino/ci763/papers/lsmtree.pdf) (1996) - O’Neil, Patrick, et al.
- [The five-minute rule ten years later, and other computer storage rules of thumb](https://dl.acm.org/doi/pdf/10.1145/271074.271094) (1997) - Gray, Jim, and Goetz Graefe.
- [The Five Minute Rule 20 Years Later and How Flash Memory Changes the Rules](https://dl.acm.org/doi/pdf/10.1145/1363189.1363198) (2008) - Graefe, Goetz.
- [A Comparison of Fractal Trees to Log-Structured Merge (LSM) Trees](http://www.pandademo.com/wp-content/uploads/2017/12/A-Comparison-of-Fractal-Trees-to-Log-Structured-Merge-LSM-Trees.pdf) (2014) - Kuszmaul, Bradley C.
- [Design Tradeoffs of Data Access Methods](https://dl.acm.org/doi/pdf/10.1145/2882903.2912569) (2016) - Athanassoulis, Manos, and Stratos Idreos.
- [Designing Access Methods: The RUM Conjecture](https://stratos.seas.harvard.edu/sites/scholar.harvard.edu/files/stratos/files/rum.pdf) (2016) - Athanassoulis, Manos, et al.
- [The five minute rule thirty years later and its impact on the storage hierarchy](https://infoscience.epfl.ch/record/230398/files/adms-talk.pdf) (2017) - Appuswamy, Raja, et al.
- [WiscKey: Separating Keys from Values in SSD-conscious Storage](https://dl.acm.org/doi/pdf/10.1145/3033273) (2017) - Lu, Lanyue, et al.
- [Managing Non-Volatile Memory in Database Systems](https://dl.acm.org/doi/pdf/10.1145/3183713.3196897) (2018) - van Renen, Alexander, et al.
- [LeanStore: In-Memory Data Management Beyond Main Memory](https://15721.courses.cs.cmu.edu/spring2020/papers/23-largethanmemory/leis-icde2018.pdf) (2018) - Leis, Viktor, et al.
- [The Case for Learned Index Structures](https://dl.acm.org/doi/pdf/10.1145/3183713.3196909) (2018) - Kraska, Tim, et al.
- [LSM-based Storage Techniques: A Survey](https://arxiv.org/pdf/1812.07527) (2019) - Luo, Chen, and Michael J. Carey.
- [Learning Multi-dimensional Indexes](https://dl.acm.org/doi/pdf/10.1145/3318464.3380579) (2019) - Nathan, Vikram, et al.
- [Umbra: A Disk-Based System with In-Memory Performance](https://db.in.tum.de/~freitag/papers/p29-neumann-cidr20.pdf) (2020) - Neumann, Thomas, and Michael J. Freitag.
- [XIndex: A Scalable Learned Index for Multicore Data Storage](https://dl.acm.org/doi/pdf/10.1145/3332466.3374547) (2020) - Tang, Chuzhe, et al.
- [The PGM-index: a fully-dynamic compressed learned index with provable worst-case bounds](https://dl.acm.org/doi/pdf/10.14778/3389133.3389135) (2020) - Ferragina, Paolo, and Giorgio Vinciguerra.
- [From WiscKey to Bourbon: A Learned Index for Log-Structured Merge Trees](https://www.usenix.org/system/files/osdi20-dai_0.pdf) (2020) - Dai, Yifan, et al.
- [CaaS-LSM: Compaction-as-a-Service for LSM-based Key-Value Stores in Storage Disaggregated Infrastructure](https://qiaolin-yu.github.io/pubs/V2mod124-yu.pdf) (2024) - Yu, Qiaolin et al.

### Transaction

- [The Notions of Consistency and Predicate Locks in a Database System](https://dl.acm.org/doi/pdf/10.1145/360363.360369) (1976) - Eswaran, Kapali P., et al.
- [Concurrency Control in Distributed Database Systems](https://dl.acm.org/doi/pdf/10.1145/356842.356846) (1981) - Bernstein, Philip A., and Nathan Goodman.
- [On Optimistic Methods for Concurrency Control](https://dl.acm.org/doi/pdf/10.1145/319566.319567) (1981) - Kung, Hsiang-Tsung, and John T. Robinson.
- [Principles of transaction-oriented database recovery](https://dl.acm.org/doi/10.1145/289.291) (1983) - Haerder, Theo, and Andreas Reuter.
- [Multiversion Concurrency Control - Theory and Algorithms](https://dl.acm.org/doi/pdf/10.1145/319996.319998) (1983) - Bernstein, Philip A., and Nathan Goodman.
- [ARIES: A transaction recovery method supporting fine-granularity locking and partial rollbacks using write-ahead logging](https://dl.acm.org/doi/pdf/10.1145/128765.128770) (1992) - Mohan C, Haderle D, Lindsay B, et al.
- [A Critique of ANSI SQL Isolation Levels](https://dl.acm.org/doi/pdf/10.1145/568271.223785) (1995) - Berenson, Hal, et al.
- [Generalized Isolation Level Definitions](https://pmg.csail.mit.edu/papers/icde00.pdf) (2000) - Adya, Atul, Barbara Liskov, and Patrick O'Neil.
- [Serializable Snapshot Isolation in PostgreSQL](https://arxiv.org/pdf/1208.4179.pdf,) (2012) - Ports, Dan RK, and Kevin Grittner.
- [Calvin: Fast Distributed Transactions for Partitioned Database Systems](https://dl.acm.org/doi/pdf/10.1145/2213836.2213838) (2012) - Thomson, Alexander, et al.
- [MaaT: effective and scalable coordination of distributed transactions in the cloud](https://dl.acm.org/doi/pdf/10.14778/2732269.2732270) (2014) - Mahmoud, Hatem A., et al.
- [Staring into the Abyss: An Evaluation of Concurrency Control with One Thousand Cores](https://dspace.mit.edu/bitstream/handle/1721.1/100022/Devadas_Staring%20into.pdf?sequence=1&isAllowed=y) (2014) - Yu, Xiangyao, et al.
- [An Evaluation of the Advantages and Disadvantages of Deterministic Database Systems](https://dl.acm.org/doi/pdf/10.14778/2732951.2732955) (2014) - Ren, Kun, Alexander Thomson, and Daniel J. Abadi.
- [Fast Serializable Multi-Version Concurrency Control for Main-Memory Database Systems](https://dl.acm.org/doi/pdf/10.1145/2723372.2749436) (2015) - Neumann, Thomas, Tobias Mühlbauer, and Alfons Kemper.
- [An Empirical Evaluation of In-Memory Multi-Version Concurrency Control](https://dl.acm.org/doi/pdf/10.14778/3067421.3067427) (2017) - Wu, Yingjun, et al.
- [An Evaluation of Distributed Concurrency Control](https://dl.acm.org/doi/pdf/10.14778/3055540.3055548) (2017) - Harding, Rachael, et al.
- [Scalable Garbage Collection for In-Memory MVCC Systems](https://dl.acm.org/doi/pdf/10.14778/3364324.3364328) (2019) - Böttcher, Jan, et al.

### Scheduling

- [Automated Demand-driven Resource Scaling in Relational Database-as-a-Service](https://dl.acm.org/doi/pdf/10.1145/2882903.2903733) (2016) - Das, Sudipto, et al.
- [Autoscaling Tiered Cloud Storage in Anna](https://dl.acm.org/doi/pdf/10.14778/3311880.3311881) (2019) - Wu, Chenggang, Vikram Sreekanti, and Joseph M. Hellerstein.
- [Adaptive HTAP through Elastic Resource Scheduling](https://dl.acm.org/doi/pdf/10.1145/3318464.3389783) (2020) - Raza, Aunn, et al.
- [MorphoSys: Automatic Physical Design Metamorphosis for Distributed Database Systems](https://dl.acm.org/doi/pdf/10.14778/3424573.3424578) (2020) - Abebe, Michael, Brad Glasbergen, and Khuzaima Daudjee.

## Miscellaneous

### Workload

- [TPC-H Analyzed: Hidden Messages and Lessons Learned from an Influential Benchmark](https://www.researchgate.net/profile/Peter-Boncz/publication/291257517_TPC-H_Analyzed_Hidden_Messages_and_Lessons_Learned_from_an_Influential_Benchmark/links/5852dbf708ae95fd8e1d749b/TPC-H-Analyzed-Hidden-Messages-and-Lessons-Learned-from-an-Influential-Benchmark.pdff) (2013) - Boncz, Peter, Thomas Neumann, and Orri Erling.
- [Quantifying TPCH Choke Points and Their Optimizations](https://dl.acm.org/doi/pdf/10.14778/3389133.3389138) (2020) - Dreseler, Markus, et al.

### Network

- [The End of Slow Networks: It's Time for a Redesign](https://arxiv.org/pdf/1504.01048) (2015) - Binnig, Carsten, et al.
- [Accelerating Relational Databases by Leveraging Remote Memory and RDMA](https://dl.acm.org/doi/pdf/10.1145/2882903.2882949) (2016) - Li, Feng, et al.
- [Don't Hold My Data Hostage: A Case for Client Protocol Redesign](https://dl.acm.org/doi/pdf/10.14778/3115404.3115408) (2017) - Raasveldt, Mark, and Hannes Mühleisen.

### Quality

- [Testing the Accuracy of Query Optimizers](https://dl.acm.org/doi/pdf/10.1145/2304510.2304525) (2012) - Gu, Zhongxian, Mohamed A. Soliman, and Florian M.

### Diagnosis and Tuning

- [Automatic SQL Tuning in Oracle 10g](http://www.vldb.org/conf/2004/IND4P2.PDF) (2004) - Dageville B, Das D, Dias K, et al.
- [Automatic Performance Diagnosis and Tuning in Oracle](https://www.cidrdb.org/cidr2005/papers/P07.pdf) (2005) - Dias K, Ramacher M, Shaft U, et al.

