import { FormEvent } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { Calendar, Tag } from 'lucide-react';
import { Modal } from '../../../../components/modal';
import { FormField } from '../../../../components/form-field';
import { Button } from '../../../../components/button';
import { api } from '../../../../lib/axios';

interface CreateActivityModalProps {
  closeCreateActivityModal: () => void;
}

export function CreateActivityModal({
  closeCreateActivityModal,
}: CreateActivityModalProps) {
  const { tripId } = useParams();
  const navigate = useNavigate();

  async function createActivity(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    const data = new FormData(event.currentTarget);
    const title = data.get('title')?.toString();
    const occurs_at = data.get('occurs_at')?.toString();
    await api.post(`/trips/${tripId}/activities`, {
      title,
      occurs_at,
    });
    navigate(0);
  }

  return (
    <Modal handleClose={closeCreateActivityModal}>
      <Modal.Header>
        <Modal.Title>Cadastrar atividade</Modal.Title>

        <Modal.SubTitle>
          Todos convidados podem visualizar as atividades.
        </Modal.SubTitle>
      </Modal.Header>

      <Modal.Content>
        <form onSubmit={createActivity} className="space-y-3">
          <FormField>
            <Tag className="text-zinc-400 size-5" />

            <FormField.Input name="title" placeholder="Qual a atividade?" />
          </FormField>

          <FormField>
            <Calendar className="text-zinc-400 size-5" />

            <FormField.Input
              type="datetime-local"
              name="occurs_at"
              placeholder="Data e horário da atividade"
            />
          </FormField>

          <Button variant="primary" size="full" type="submit">
            <span>Salvar atividade</span>
          </Button>
        </form>
      </Modal.Content>
    </Modal>
  );
}