import { FC } from 'react';
import { Transaction } from './types';
import { CloseButton, HStack, Td, Text, Tr } from '@chakra-ui/react';
import EditableField from './EditableField';
import { isDate, isNotEmpty, isNumber } from './utils';

interface TransactionRowProps {
  transaction: Transaction;
  deleteTransaction: (id: number) => void;
  updateTransaction: (transaction: Transaction) => void;
}

const TransactionRow: FC<TransactionRowProps> = ({
  transaction,
  deleteTransaction,
  updateTransaction,
}) => {
  const { id, date, amount, description, category } = transaction;

  const update = (prop: keyof Transaction, val: string) => {
    const newTransaction = { ...transaction, [prop]: val };
    updateTransaction(newTransaction);
  };

  return (
    <Tr key={id}>
      <Td w="30px">
        <EditableField
          initialValue={date}
          onSubmit={(val) => update('date', val)}
          type="date"
          validate={isDate}
        />
      </Td>
      <Td w="30px" pl="30px" pr={0}>
        <HStack>
          <Text m={-2} p={0}>
            $
          </Text>
          <EditableField
            initialValue={`${amount}`}
            onSubmit={(val) => update('amount', val)}
            placeholder="00.00"
            validate={isNumber}
          />
        </HStack>
      </Td>
      <Td minW="30px" maxW="30px" whiteSpace="normal">
        <EditableField
          initialValue={description}
          onSubmit={(val) => update('description', val)}
          placeholder="description"
          validate={isNotEmpty}
        />
      </Td>
      <Td w="30px" px={0}>
        <EditableField
          initialValue={category}
          onSubmit={(val) => update('category', val)}
          placeholder="category"
          type="select"
        />
      </Td>
      <Td>
        <CloseButton
          color="red"
          alignSelf="center"
          onClick={() => deleteTransaction(id)}
        />
      </Td>
    </Tr>
  );
};

export default TransactionRow;
