# Generalized listing deposit mechanism for collateral

The $i$th highest bid price is denoted as $p_i$.

$n$ is a parameter determined by the NFT exhibitor. The default value is $n=10$.

In this system, bidders have to deposit their balance in proportion to the bidding price.

The deposit amount of the bidder of the $i$th highest bid price is denoted as $d_i$.

We define that

$$
d_i = \frac{1}{n} p_i
$$

The maximum value of using NFT as a collateral will be denoted as $q$.

In this system, we define that

$$
q = \sum_{i=1}^n d_i = d_1 + \cdots + d_n = \frac{1}{n} \sum_{i=1}^n p_i
$$

It means that, the maximum value of using NFT as a collateral is, an average of top $n$ highest bid prices.

### Bid cancellation fee

If the bidder want to cancel his bid, some fee may occur.

The borrowed value by NFT exhibitor is denoted as $b$.

The cancellation fee of the bidder of the $i$th highest bid price is

$$
\begin{cases}
\max\{d_i - (q - b), 0\} & \ \text{if} \ i \le n \\
0 & \ \text{if} \ i > n
\end{cases}
$$

It means that, if the bidder is in bidders of top $n$ highest bid price, their deposit may be forfeited.

### Settlement with exhibitor’s decision

The 1st highest bidder have to pay $p_1 - d_1$ during the period of payment.

If he doesn’t do so, his deposit $d_1$ will be forfeited and he will be removed from bidders.

### Liquidation

If $d_1 \ge q$, it means that, the deposit amount of the bidder of the 1st highest bid price is greater than or equal to $q$, the 1st highest bidder will receive the NFT and the charge $d_1 - q$.

In other cases, the procedure below will be iterated for $n$ times. In the $i$th iteration,

The $i$th highest bidder have to pay $p_i - d_i$ during the period of payment.

If he does so, he will receive the NFT and the iteration will be stopped.

If he doesn’t do so, his deposit $d_i$ will be forfeited and the iteration continues to the next $i$.
