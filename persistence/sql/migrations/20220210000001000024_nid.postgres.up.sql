-- Migration generated by the command below; DO NOT EDIT.
-- hydra:generate hydra migrate gen

ALTER TABLE hydra_oauth2_flow ALTER nid SET NOT NULL;