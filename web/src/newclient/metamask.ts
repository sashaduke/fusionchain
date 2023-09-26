import { TxContext } from "@evmos/transactions";
import { broadcastTransaction, buildTransaction, fetchAccount, signTransactionMetamask } from ".";
import { chain } from "../keplr";
import { Message } from "@bufbuild/protobuf";
import { hashMessage } from '@ethersproject/hash'
import {
  computePublicKey,
  recoverPublicKey,
} from '@ethersproject/signing-key'

export async function metamaskBuildAndBroadcast(msgs: Message<any>[]) {
  const accounts = await window?.ethereum?.request({
    method: 'eth_requestAccounts',
  })

  // Handle errors if MetaMask fails to return any accounts.
  const message = 'Verify Public Key'

  const signature = await window?.ethereum?.request({
    method: 'personal_sign',
    params: [message, accounts[0], ''],
  })

  // Compress the key, since the client expects
  // public keys to be compressed.
  const uncompressedPk = recoverPublicKey(
    hashMessage(message),
    signature,
  )

  const hexPk = computePublicKey(uncompressedPk, true)
  const pubkey = Buffer.from(
    hexPk.replace('0x', ''), 'hex',
  ).toString('base64')

  // fetch sequence number
  const chainAccount = await fetchAccount(accounts[0].bech32Address);

  // build tx context
  const context: TxContext = {
    chain,
    sender: {
      accountAddress: chainAccount.account.base_account.address,
      sequence: parseInt(chainAccount.account.base_account.sequence),
      accountNumber: parseInt(chainAccount.account.base_account.account_number),
      pubkey,
    },
    fee: {
      amount: '200000000000000',
      denom: 'nQRDO',
      gas: '200000',
    },
    memo: "",
  }

  // 1 - build tx
  const tx = buildTransaction(context, msgs);
  // 2 - sign tx
  const signedTx = await signTransactionMetamask(context, tx);
  // 3 - broadcast tx
  const res = await broadcastTransaction(signedTx, undefined);

  return res
}

