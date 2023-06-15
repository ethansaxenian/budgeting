import { FC, useState } from 'react';
import TransactionsTable from './TransactionsTable';
import { Box, Button, HStack, Heading, Text, VStack } from '@chakra-ui/react';
import { Category, Month, Plan, Transaction, TransactionType } from './types';
import AddTransactionModal from './AddTransactionModal';
import { getTransactions, postTransaction } from './api';
import PlansTable from './PlansTable';
import { colorAmount, formatAmount } from './utils';

type MonthPageProps = {
  month: Month;
};

const MonthPage: FC<MonthPageProps> = ({ month }) => {
  const [incomeTransactions, setIncomeTransactions] = useState<Transaction[]>(
    []
  );
  const [plannedIncome, setPlannedIncome] = useState<Plan[]>([]);
  const [expenseTransactions, setExpenseTransactions] = useState<Transaction[]>(
    []
  );
  const [plannedExpenses, setPlannedExpenses] = useState<Plan[]>([]);
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

  const totalPlannedExpenses =
    Math.round(plannedExpenses.reduce((sum, e) => sum + e.amount, 0) * 100) /
    100;

  const totalIncome =
    Math.round(incomeTransactions.reduce((sum, i) => sum + i.amount, 0) * 100) /
    100;

  const totalPlannedIncome =
    Math.round(plannedIncome.reduce((sum, e) => sum + e.amount, 0) * 100) / 100;

  return (
    <VStack>
      <Heading>{month.id}</Heading>
      <Box>Starting Balance: {formatAmount(month.starting_balance)}</Box>
      <Box>
        Ending Balance:{' '}
        {formatAmount(month.starting_balance + totalIncome - totalExpenses)}
      </Box>
      <HStack>
        <Text>Amount Saved: </Text>
        <Text color={colorAmount(totalIncome - totalExpenses)}>
          {formatAmount(totalIncome - totalExpenses)}
        </Text>
      </HStack>
      <HStack alignItems="flex-start" my={50}>
        <VStack>
          <Heading size="lg">Planned Expenses</Heading>
          <HStack justifyContent="space-around">
            <VStack alignItems="flex-start">
              <Text>Planned:</Text>
              <Text>Actual:</Text>
            </VStack>
            <VStack alignItems="flex-start">
              <Text>{formatAmount(totalPlannedExpenses)}</Text>
              <Text color={colorAmount(totalExpenses)}>
                {formatAmount(totalExpenses)}
              </Text>
            </VStack>
          </HStack>
          <PlansTable
            plans={plannedExpenses}
            transactions={expenseTransactions}
            setPlans={setPlannedExpenses}
            monthId={month.id}
            type={TransactionType.Expense}
          />
        </VStack>
        <VStack>
          <Heading size="lg">Planned Income</Heading>
          <HStack justifyContent="space-around">
            <VStack alignItems="flex-start">
              <Text>Planned:</Text>
              <Text>Actual:</Text>
            </VStack>
            <VStack alignItems="flex-start">
              <Text>{formatAmount(totalPlannedIncome)}</Text>
              <Text color={colorAmount(totalIncome)}>
                {formatAmount(totalIncome)}
              </Text>
            </VStack>
          </HStack>
          <PlansTable
            plans={plannedIncome}
            transactions={incomeTransactions}
            setPlans={setPlannedIncome}
            monthId={month.id}
            type={TransactionType.Income}
          />
        </VStack>
      </HStack>

      <HStack>
        <VStack alignSelf="flex-start">
          <HStack>
            <Heading size="lg" mr={3}>
              Expenses
            </Heading>
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
            <Heading size="lg" mr={3}>
              Income
            </Heading>
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
