# Integrate Yield Farming

Introduction to integrating Yield Farming created from cosmwasm

When Yield Farming is registered in the `yield-aggregator`, a Denom is automatically sent to that address.

Then `yield-aggregator` send `FarmingUnitMsg` to Yield Farming.

Yield Farming operate the service with the received Msg based on that information.

Send the Denom back to the module with `AssetManagementAccountBankKeeper.PayBack` within the Unbonding_time.

## abstract YA work flow

```mermaid
graph TD;
  subgraph YA

    fire_Farming_Order_Event-->Process_of_selecting_YF_from_FO;
    Process_of_selecting_YF_from_FO-->send_denom_to_YF;
    send_denom_to_YF-->send_FarmingUnitMsg_to_YF;

    receive_denom_on_YA-->check_FarmingUnit_on_YA;
    check_FarmingUnit_on_YA-->update_DPR_process;
  end

  subgraph YF
    send_FarmingUnitMsg_to_YF-->receive_denom_on_YF;
    send_denom_to_YF-->receive_denom_on_YF;
    receive_denom_on_YF-->check_FarmingUnit_on_YF;
    check_FarmingUnit_on_YF-->earning_denom;
    earning_denom-->finish_unbonding_time;
    finish_unbonding_time-->payback_denom_to_YA;
    payback_denom_to_YA-->receive_denom_on_YA;
  end
```
