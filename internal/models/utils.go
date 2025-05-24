package models

func RetrieveTarefas() ([]Tarefa, error) {
	var tarefas []Tarefa
	if err := db.Find(&tarefas).Error; err != nil {
		return nil, err
	}
	return tarefas, nil
}

func RetrieveTarefaById(id int) (Tarefa, error) {
	var tarefa Tarefa
	if err := db.First(&tarefa, id).Error; err != nil {
		return Tarefa{}, err
	}
	return tarefa, nil
}

func CreateTarefa(tarefa *Tarefa) (error) {
	if err := db.Create(tarefa).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTarefa(tarefa Tarefa) (error) {
	if err := db.Save(&tarefa).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTarefa(id int) (error) {
	if err := db.Delete(&Tarefa{}, id).Error; err != nil {
		return err
	}
	return nil
}