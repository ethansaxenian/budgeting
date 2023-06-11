import {
  Box,
  Button,
  Input,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  NumberDecrementStepper,
  NumberIncrementStepper,
  NumberInput,
  NumberInputField,
  NumberInputStepper,
  Select,
  VStack,
} from '@chakra-ui/react';
import { FC, useState } from 'react';
import { Category, TransactionType } from './types';
import DatePicker from 'react-datepicker';

interface AddTransactionModalProps {
  isOpen: boolean;
  create: (
    date: Date,
    amount: number,
    description: string,
    category: Category,
    transactionType: TransactionType
  ) => void;
  onClose: () => void;
}

const AddTransactionModal: FC<AddTransactionModalProps> = ({
  isOpen,
  create,
  onClose,
}) => {
  const [date, setDate] = useState<Date>(new Date());
  const [amount, setAmount] = useState('');
  const [description, setDescription] = useState('');
  const [category, setCategory] = useState<Category>(Category.Transportation);
  const [transactionType, setTransactionType] = useState<TransactionType>(
    TransactionType.Expense
  );

  return (
    <Modal isOpen={isOpen} onClose={onClose}>
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>Add New Transaction</ModalHeader>
        <ModalCloseButton />
        <ModalBody>
          <VStack alignItems="flex-start">
            <Box
              px="12px"
              py="8px"
              border="1px solid lightgray"
              borderRadius="6px"
            >
              <DatePicker
                selected={date}
                onChange={(date) => setDate(date as Date)}
              />
            </Box>
            <NumberInput
              value={amount}
              min={0}
              step={1}
              precision={2}
              onChange={(val) => setAmount(val)}
            >
              <NumberInputField />
              <NumberInputStepper>
                <NumberIncrementStepper />
                <NumberDecrementStepper />
              </NumberInputStepper>
            </NumberInput>
            <Input
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              placeholder="description"
            />
            <Select
              placeholder="Select a Category"
              isRequired
              onChange={(e) => setCategory(e.target.value as Category)}
              value={category}
            >
              {Object.keys(Category).map((name) => (
                <option key={name} value={name}>
                  {name}
                </option>
              ))}
            </Select>
            <Select
              isRequired
              onChange={(e) =>
                setTransactionType(e.target.value as TransactionType)
              }
              value={transactionType}
            >
              {Object.keys(TransactionType).map((name) => (
                <option key={name} value={name}>
                  {name}
                </option>
              ))}
            </Select>
          </VStack>
        </ModalBody>

        <ModalFooter>
          <Button
            colorScheme="green"
            onClick={() =>
              create(
                date,
                parseFloat(amount || '0'),
                description,
                category,
                transactionType
              )
            }
          >
            Add Transaction
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};

export default AddTransactionModal;
