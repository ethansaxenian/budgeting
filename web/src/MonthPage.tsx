import { FC, useState } from 'react';
import TransactionsTable from './TransactionsTable';
import { Box, Button, HStack, Heading, VStack } from '@chakra-ui/react';
import { Category, Month, Transaction, TransactionType } from './types';
import AddTransactionModal from './AddTransactionModal';
import { getTransactions, postTransaction } from './api';

type MonthPageProps = {
  month: Month;
};

const MonthPage: FC<MonthPageProps> = ({ month }) => {
  const [incomeTransactions, setIncomeTransactions] = useState<Transaction[]>(
    []
  );
  const [expenseTransactions, setExpenseTransactions] = useState<Transaction[]>(
    []
  );
  const [addingTransaction, setAddingTransaction] = useState(false);

  const addTransaction = async (
    date: Date,
    amount: number,
    description: string,
    category: Category,
    transactionType: TransactionType
  ) => {
    await postTransaction(date, amount, description, category, transactionType);

    setAddingTransaction(false);

    const data = await getTransactions(month.id, transactionType);

    if (transactionType === TransactionType.Income) {
      setIncomeTransactions(data);
    } else {
      setExpenseTransactions(data);
    }
  };

  const totalExpenses =
    Math.round(
      expenseTransactions.reduce((sum, e) => sum + e.amount, 0) * 100
    ) / 100;

  const totalIncome =
    Math.round(incomeTransactions.reduce((sum, i) => sum + i.amount, 0) * 100) /
    100;

  return (
    <VStack>
      <Heading>{month.id}</Heading>
      <Box>Starting Balance: ${month.startingBalance}</Box>
      <Box>
        Ending Balance: ${month.startingBalance + totalIncome - totalExpenses}
      </Box>
      <Box>Amount Saved: ${totalIncome - totalExpenses}</Box>
      <Box>Actual Income: ${totalIncome}</Box>

      <HStack>
        <VStack alignSelf="flex-start">
          <HStack>
            <Heading>Expenses</Heading>
            <Button
              colorScheme="green"
              onClick={() => setAddingTransaction(true)}
            >
              Add
            </Button>
          </HStack>
          <TransactionsTable
            monthId={month.id}
            type={TransactionType.Expense}
            transactions={expenseTransactions}
            setTransactions={setExpenseTransactions}
          />
        </VStack>

        <VStack alignSelf="flex-start">
          <HStack>
            <Heading>Income</Heading>
            <Button
              colorScheme="green"
              onClick={() => setAddingTransaction(true)}
            >
              Add
            </Button>
          </HStack>
          <TransactionsTable
            monthId={month.id}
            type={TransactionType.Income}
            transactions={incomeTransactions}
            setTransactions={setIncomeTransactions}
          />
        </VStack>
      </HStack>
      <AddTransactionModal
        onClose={() => setAddingTransaction(false)}
        isOpen={addingTransaction}
        create={addTransaction}
      />
    </VStack>
  );
};

export default MonthPage;
