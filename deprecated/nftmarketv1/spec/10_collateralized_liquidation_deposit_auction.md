# Collateral Liquidation Deposit Auction

## Definition

- $i \in I$: index of bids
- $n = |I|$: number of bids
- $\{p_i\}_{i \in I}$: the price of $i$ th bid
- $\{d_i\}_{i \in I}$: the deposit amount of $i$ th bid
- $\{r_i\}_{i \in I}$: the interest rate of $i$ th bid
- $\{t_i\}_{i \in I}$: the expiration date of $i$ th bid
- $q = \frac{1}{n} \sum_{i \in I} p_i$
- $s = \sum_{i \in I} d_i$: means the amount which lister can borrow with NFT as collateral
- $\{a_i\}_{i \in I}$: means the amount borrowed from $i$ th bid deposit
- $b = \sum_{i \in I} a_i$
- $i_p(j)$: means the index of the $j$ th highest price bid
- $i_d(j)$: means the index of the $j$ th highest deposit amount bid
- $i_r(j)$: means the index of the $j$ th lowest interest rate bid
- $i_t(j)$: means the index of the $j$ th farthest expiration date bid
- $c$: minimum deposit rate

## State transition

### New bid

When $(p_{\text{new}}, d_{\text{new}}, r_{\text{new}}, t_{\text{new}})$ will be added to the set of bids, the new bids sequence will be

- $I' = I \cup \{n+1\}$
- $n' = n + 1$
- $\{p_i'\}_{i \in I'} = \{p_i\}_{i \in I} \cup \{p_{\text{new}}\}$
- $\{d_i'\}_{i \in I'} = \{d_i\}_{i \in I} \cup \{d_{\text{new}}\}$
- $\{r_i'\}_{i \in I'} = \{r_i\}_{i \in I} \cup \{r_{\text{new}}\}$
- $\{t_i'\}_{i \in I'} = \{t_i\}_{i \in I} \cup \{t_{\text{new}}\}$
- $q' = \frac{1}{n'} \sum_{i \in I'} p_i'$
- $s' = \sum_{i \in I'} d_i'$

where the prime means the next state.

The constraint of $d_{n+1}'$ is

$$
  c p_{n+1}' \le d_{n+1}' \le q' - s
$$

In easy expression, it means

- $c p_{n+1}' \le d_{n+1}'$
- $s' = s + d_{n+1}' \le q'$

where $c$ means minimum deposit rate.

### New borrowing

$a_i$ must follow the constraint

- $a_i \le d_i$
  - Trivially, the following inequation must be satisfied: $b \le s$
- $a_{i_r(j+1)} = 0 \ \text{if} \ a_{i_r(j)} < d_{i_r(j)}$
  - It means that deposited amount must be consumed (used for lending resource) in ascending order of interest rates.

## User Interface guideline

### Deposited amount graph

- Horizontal axis expresses the time.
- Vertical axis expresses the deposited amount.
- The lower $i_t^{-1}$ of the deposit (the farther expiration date), the lower the deposit will be depicted in the graph as a rectangle.

### Borrowed amount graph

- Horizontal axis expresses the time.
- Vertical axis expresses the borrowed amount.
- The lower $i_r^{-1}$ of the deposit (the lower interest rate), the lower the deposit will be depicted in the graph as a rectangle.
