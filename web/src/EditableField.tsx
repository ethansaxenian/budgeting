import {
  Editable,
  EditableInput,
  EditablePreview,
  Select,
} from '@chakra-ui/react';
import { FC, useEffect, useState } from 'react';
import { Category } from './types';

interface EditableFieldProps {
  initialValue: string | number;
  onSubmit: (value: string) => void;
  placeholder?: string;
  type?: string;
  validate?: (value: string) => boolean;
}

const EditableField: FC<EditableFieldProps> = ({
  initialValue,
  onSubmit,
  placeholder,
  type,
  validate,
}) => {
  const [value, setValue] = useState<string>(initialValue as string);

  useEffect(() => {
    setValue(initialValue as string);
  }, [initialValue]);

  const submit = (value: string) => {
    if (!validate || validate(value)) {
      onSubmit(value);
    } else {
      setValue(initialValue as string);
    }
  };

  return type === 'select' ? (
    <Select
      isRequired
      onChange={(e) => {
        setValue(e.target.value);
        submit(e.target.value);
      }}
      value={value}
      size="xs"
      w="100%"
      variant="unstyled"
    >
      {Object.keys(Category).map((name) => (
        <option key={name} value={name}>
          {name}
        </option>
      ))}
    </Select>
  ) : (
    <Editable
      defaultValue={initialValue as string}
      value={value}
      onChange={setValue}
      onSubmit={submit}
      placeholder={placeholder || ''}
    >
      <EditablePreview />
      <EditableInput type={type || 'input'} />
    </Editable>
  );
};

export default EditableField;
