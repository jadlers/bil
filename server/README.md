# Server written with DDD architecture

This project is a smaller (at least initially) one in which I'm attempting to
learn more about design driven development.

## Important about aggregates

> An aggregate is a cluster of associated objects that we treat as a unit for
> the purpose of data changes. Each aggregate has a root and a boundary. The
> boundary defines what is inside the aggregate. The root is a single, specific
> entity contained in the aggregate. The root is the only member of the
> aggregate that outside objects are allowed to hold reference to, although
> objects within the boundary may hold references to each other. Entities other
> than the root have local identity, but that identity needs to be
> distinguishable only within the aggregate, because no outside object can ever
> see it out of the context of the root entity.
>
>   -- https://www.eventstore.com/event-sourcing#Write-model

## Four rules of thumb

> 1. Protect business invariants inside Aggregate boundaries
> 2. Design small aggregates
> 3. Reference other aggregates by identity only
> 4. Update other aggregates using eventual consistency only
>
>   -- https://www.youtube.com/watch?v=Xf_aLAK1RfE

## Model true invariants in consistency

> When trying to discover the aggregates in a bounded context, we must
> understand the model's true invariants. Only with that knowledge can we
> determine which objects should be clustered into a given aggregate.
>
> An invariant is a business rule that must always be consistent. There are
> different kinds of consistency. One is transactional, which is considered
> immediate and atomic. There is also eventual consistency. When discussing
> invariants, we are referring to transactional consistency.
>
>   -- Part I: https://www.dddcommunity.org/library/vernon_2011/
