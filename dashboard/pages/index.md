---
title: Тихонов Н.В инжиниринг данных
---

```sql proms
    select id, promm from localdb.proms ; 
```


<Dropdown
data={proms}
name=prom
value=promm
/>

<Value ${inputs.prom.values} />

```sql promID 
  select id from localdb.proms where promm = '${inputs.prom.value}'
```

```sql counts
    select * from localdb.counts where prom_id  =  ${promID}; 
```

```sql like
    select * from localdb.likes where prom_id  =  ${promID}; 
```

```sql view
    select * from localdb.views where prom_id  =  ${promID}; 
```


```sql passion
    select * from localdb.passions where prom_id  =  ${promID}; 
```

```sql comment
    select * from localdb.comments where prom_id  =  ${promID} ;
```


```sql activities
    select * from localdb.activities where prom_id = ${promID} ;
```


<BarChart
data={counts}
x=year
y=count
xAxisTitle="Year"
yAxisTitle="Count"
title="count"
/>

<BarChart
data={view}
x=year
y=count
xAxisTitle="Year"
yAxisTitle="Count"
title="views"
/>

<BarChart
data={like}
x=year
y=count
xAxisTitle="Year"
yAxisTitle="Count"
title="likes"
/>

<BarChart
data={comment}
x=year
y=count
xAxisTitle="Year"
yAxisTitle="Count"
title="comments"
/>

<BarChart
data={passion}
x=year
y=proc
xAxisTitle="Year"
yAxisTitle="Proc"
title="passions"
/>

<BarChart
data={activities}
x=year
y=proc
xAxisTitle="Year"
yAxisTitle="Proc"
title="activities"
/>
