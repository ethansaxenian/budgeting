import { FC, useEffect } from 'react';
import { deleteTransaction, getTransactions } from './api';
import {
  CloseButton,
  Table,
  TableContainer,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
} from '@chakra-ui/react';
import { Transaction, TransactionType } from './types';
import { strToDate } from './utils';

interface TransactionsTableProps {
  monthId: string;
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

  const sortedTransactions = transactions.sort(
    (a, b) => strToDate(b.date).getDate() - strToDate(a.date).getDate()
  );

  return (
    <TableContainer>
      <Table fontSize="12px">
        <Thead>
          <Tr>
            <Th>Date</Th>
            <Th>Amount</Th>
            <Th>Description</Th>
            <Th>Category</Th>
            <Th />
          </Tr>
        </Thead>
        <Tbody>
          {sortedTransactions.map((transaction) => (
            <Tr key={transaction.id}>
              <Td w="30px">{transaction.date}</Td>
              <Td w="30px">{transaction.amount}</Td>
              <Td maxW="30px" whiteSpace="normal">
                {transaction.description}
              </Td>
              <Td w="30px">{transaction.category}</Td>
              <Td>
                <CloseButton
                  color="red"
                  alignSelf="center"
                  onClick={() => deleteTransactionWithId(transaction.id)}
                />
              </Td>
            </Tr>
          ))}
        </Tbody>
      </Table>
    </TableContainer>
  );
};

export default TransactionsTable;
