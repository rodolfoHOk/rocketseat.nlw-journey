import { DateRange, DayPicker } from 'react-day-picker';
import { Modal } from '../../../../components/modal';

interface RangeDatePickerModalProps {
  tripStartAndEndDates: DateRange | undefined;
  closeDatePicker: () => void;
  setTripStartAndEndDates: (dates: DateRange | undefined) => void;
}

export function RangeDatePickerModal({
  tripStartAndEndDates,
  closeDatePicker,
  setTripStartAndEndDates,
}: RangeDatePickerModalProps) {
  return (
    <Modal width="fit" handleClose={closeDatePicker}>
      <Modal.Header>
        <Modal.Title>Selecione a data</Modal.Title>
      </Modal.Header>

      <Modal.Content>
        <DayPicker
          mode="range"
          disabled={{ before: new Date() }}
          selected={tripStartAndEndDates}
          onSelect={setTripStartAndEndDates}
        />
      </Modal.Content>
    </Modal>
  );
}
