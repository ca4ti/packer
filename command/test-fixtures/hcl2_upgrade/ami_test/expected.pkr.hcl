
source "amazon-ebs" "autogenerated_1" {
  run_tags = {
    SourceAMI             = "{{ .SourceAMI }}"
    SourceAMICreationDate = "{{ .SourceAMICreationDate }}"
    SourceAMIName         = "{{ .SourceAMIName }}"
    SourceAMIOwner        = "{{ .SourceAMIOwner }}"
    SourceAMIOwnerName    = "{{ .SourceAMIOwnerName }}"
    SourceAMITags         = "{{ .SourceAMITags.TagName }}"
  }
  tags = {
    SourceAMI             = "{{ .SourceAMI }}"
    SourceAMICreationDate = "{{ .SourceAMICreationDate }}"
    SourceAMIName         = "{{ .SourceAMIName }}"
    SourceAMIOwner        = "{{ .SourceAMIOwner }}"
    SourceAMIOwnerName    = "{{ .SourceAMIOwnerName }}"
    SourceAMITags         = "{{ .SourceAMITags.TagName }}"
  }
}

build {
  sources = ["source.amazon-ebs.autogenerated_1"]

}
