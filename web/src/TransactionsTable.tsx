import { FC, useEffect, useState } from 'react';
import { deleteTransaction, getTransactions, putTransaction } from './api';
import { Table, TableContainer, Tbody, Th, Thead, Tr } from '@chakra-ui/react';
import { Transaction, TransactionType } from './types';
import { sortTransactions } from './utils';
import TransactionRow from './TransactionRow';
import { ChevronDownIcon, ChevronUpIcon } from '@chakra-ui/icons';

interface TransactionsTableProps {
  monthId: number;
  type: TransactionType;
  transactions: Transaction[];
  setTransactions: (transactions: Transaction[]) => void;
}

const TransactionsTable: FC<TransactionsTableProps> = ({
  monthId,
  type,
  transactions,
  setTransactions,
}) => {
  const [sort, setSort] = useState<[keyof Transaction, boolean]>([
    'date',
    false,
  ]);

  useEffect(() => {
    const fetchTransactions = async () => {
      const data = await getTransactions(monthId, type);
      setTransactions(data);
    };

    fetchTransactions();
  }, [monthId, type, setTransactions]);

  const deleteTransactionWithId = async (id: number) => {
    await deleteTransaction(id);
    const data = await getTransactions(monthId, type);
    setTransactions(data);
  };

  const updateTransaction = async (transaction: Transaction) => {
    await putTransaction(transaction);
    const data = await getTransactions(monthId, type);
    setTransactions(data);
  };

  const updateSort = (field: keyof Transaction) => {
    if (sort[0] === field) {
      setSort([sort[0], !sort[1]]);
    } else {
      setSort([field, true]);
    }
  };

  const sortedTransactions = sortTransactions(transactions, ...sort);

  return (
    <TableContainer>
      <Table fontSize="12px">
        <Thead>
          <Tr>
            <Th onClick={() => updateSort('date')} cursor="pointer">
              Date{' '}
              {sort[0] === 'date' ? (
                sort[1] ? (
                  <ChevronUpIcon />
                ) : (
                  <ChevronDownIcon />
                )
              ) : null}
            </Th>
            <Th onClick={() => updateSort('amount')} cursor="pointer">
              Amount{' '}
              {sort[0] === 'amount' ? (
                sort[1] ? (
                  <ChevronUpIcon />
                ) : (
                  <ChevronDownIcon />
                )
              ) : null}
            </Th>
            <Th>Description</Th>
            <Th>Category</Th>
            <Th />
          </Tr>
        </Thead>
        <Tbody>
          {sortedTransactions.map((transaction) => (
            <TransactionRow
              key={transaction.id}
              transaction={transaction}
              deleteTransaction={deleteTransactionWithId}
              updateTransaction={updateTransaction}
            />
          ))}
        </Tbody>
      </Table>
    </TableContainer>
  );
};

export default TransactionsTable;
